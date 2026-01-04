import { ref as storageRef, uploadBytesResumable, getDownloadURL, deleteObject } from 'firebase/storage'
import { storage } from '@/plugins/firebase'

/**
 * Upload a file to Firebase Storage
 * @param {File} file - The file to upload
 * @param {string} path - The storage path (e.g., 'products/image.jpg')
 * @param {Function} onProgress - Callback for upload progress (0-100)
 * @returns {Promise<string>} - Download URL of uploaded file
 */
export async function uploadFile(file, path, onProgress = null) {
  if (!file) {
    throw new Error('No file provided')
  }

  // Create a storage reference
  const fileRef = storageRef(storage, path)

  // Upload file
  const uploadTask = uploadBytesResumable(fileRef, file)

  return new Promise((resolve, reject) => {
    uploadTask.on(
      'state_changed',
      (snapshot) => {
        // Track progress
        const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100
        if (onProgress) {
          onProgress(progress)
        }
      },
      (error) => {
        // Handle errors
        console.error('Upload error:', error)
        reject(error)
      },
      async () => {
        // Upload completed successfully, get download URL
        try {
          const downloadURL = await getDownloadURL(uploadTask.snapshot.ref)
          resolve(downloadURL)
        } catch (error) {
          reject(error)
        }
      }
    )
  })
}

/**
 * Upload product image
 * @param {File} file - Image file
 * @param {string} productId - Product ID (optional, generates unique name if not provided)
 * @param {Function} onProgress - Progress callback
 * @returns {Promise<string>} - Download URL
 */
export async function uploadProductImage(file, productId = null, onProgress = null) {
  // Validate file type
  const validTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/webp']
  if (!validTypes.includes(file.type)) {
    throw new Error('Invalid file type. Only JPEG, PNG, and WebP images are allowed.')
  }

  // Validate file size (max 5MB)
  const maxSize = 5 * 1024 * 1024 // 5MB
  if (file.size > maxSize) {
    throw new Error('File size too large. Maximum size is 5MB.')
  }

  // Generate unique filename
  const timestamp = Date.now()
  const randomString = Math.random().toString(36).substring(2, 8)
  const extension = file.name.split('.').pop()
  const filename = productId
    ? `${productId}_${timestamp}.${extension}`
    : `product_${timestamp}_${randomString}.${extension}`

  // Upload to products folder
  const path = `products/${filename}`

  return uploadFile(file, path, onProgress)
}

/**
 * Delete a file from Firebase Storage
 * @param {string} url - Download URL of the file
 * @returns {Promise<void>}
 */
export async function deleteFile(url) {
  try {
    // Extract path from URL
    const fileRef = storageRef(storage, url)
    await deleteObject(fileRef)
  } catch (error) {
    console.error('Delete error:', error)
    throw error
  }
}

/**
 * Delete product image by URL
 * @param {string} imageUrl - The image URL to delete
 * @returns {Promise<void>}
 */
export async function deleteProductImage(imageUrl) {
  if (!imageUrl || !imageUrl.includes('firebasestorage.googleapis.com')) {
    console.warn('Not a Firebase Storage URL, skipping delete')
    return
  }

  try {
    // Extract path from URL
    const decodedUrl = decodeURIComponent(imageUrl)
    const pathMatch = decodedUrl.match(/\/o\/(.+?)\?/)

    if (pathMatch && pathMatch[1]) {
      const path = pathMatch[1]
      const fileRef = storageRef(storage, path)
      await deleteObject(fileRef)
    }
  } catch (error) {
    console.error('Error deleting product image:', error)
  }
}

/**
 * Compress image before upload (client-side)
 * @param {File} file - Image file
 * @param {number} maxWidth - Max width (default 1200)
 * @param {number} quality - Quality 0-1 (default 0.8)
 * @returns {Promise<Blob>}
 */
export function compressImage(file, maxWidth = 1200, quality = 0.8) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)

    reader.onload = (event) => {
      const img = new Image()
      img.src = event.target.result

      img.onload = () => {
        const canvas = document.createElement('canvas')
        let width = img.width
        let height = img.height

        if (width > maxWidth) {
          height = (height * maxWidth) / width
          width = maxWidth
        }

        canvas.width = width
        canvas.height = height

        const ctx = canvas.getContext('2d')
        ctx.drawImage(img, 0, 0, width, height)

        canvas.toBlob(
          (blob) => {
            resolve(blob)
          },
          file.type,
          quality
        )
      }

      img.onerror = reject
    }

    reader.onerror = reject
  })
}
