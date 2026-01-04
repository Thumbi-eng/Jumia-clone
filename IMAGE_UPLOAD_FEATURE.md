# Product Image Upload Feature

## Overview

The admin panel now supports uploading product images directly to Firebase Storage instead of just using URL strings. This provides a professional, production-ready image management system.

## Features

✅ **Drag & Drop Upload**

- Click to select files
- Drag and drop images directly
- Real-time preview

✅ **Image Validation**

- File type: JPEG, PNG, WebP only
- Maximum size: 5MB
- Client-side validation before upload

✅ **Image Compression**

- Automatic compression before upload
- Max dimensions: 1200px
- 80% JPEG quality
- Reduces file size and storage costs

✅ **Upload Progress**

- Visual progress bar
- Percentage display
- Loading state indicators

✅ **Flexible Input**

- Upload files from computer
- Or paste image URL manually
- Preview for both methods

✅ **Image Management**

- Remove uploaded images
- Re-upload if needed
- Delete from Firebase Storage

## How to Use

### 1. Set Up Firebase (One-time)

Follow the instructions in [FIREBASE_SETUP.md](./FIREBASE_SETUP.md):

1. Create Firebase project
2. Enable Storage
3. Copy configuration
4. Update `.env` file
5. Restart dev server

### 2. Add Product with Image

1. Navigate to **Admin > Add Products** (`/admin/products/add`)
2. Fill in product details (name, description, price, etc.)
3. **Upload Image**:
   - Click "Choose File" button
   - Or drag image onto the upload area
   - Wait for preview to appear
   - Click "Upload Image" button
   - Watch progress bar complete
4. Click "Add Product" to save

### 3. Alternative: Use Image URL

If you don't want to upload:

1. Host image elsewhere (Imgur, Cloudinary, etc.)
2. Copy the direct image URL
3. Paste into "Or enter image URL" field
4. Submit form

## File Structure

```
src/
├── plugins/
│   └── firebase.js          # Firebase initialization
├── utils/
│   └── storage.js           # Upload/delete/compress utilities
└── views/
    └── admin/
        └── AddProducts.vue  # Image upload component
```

## Code Example

### Upload Image

```javascript
import { uploadProductImage } from "@/utils/storage";

const imageFile = ref(null);
const uploadProgress = ref(0);

async function uploadImage() {
  const downloadURL = await uploadProductImage(imageFile.value, (progress) => {
    uploadProgress.value = progress;
  });

  // Use downloadURL in product data
  productData.image_url = downloadURL;
}
```

### Delete Image

```javascript
import { deleteProductImage } from "@/utils/storage";

async function removeImage() {
  const imageUrl = "https://firebasestorage.googleapis.com/...";
  await deleteProductImage(imageUrl);
}
```

### Compress Image

```javascript
import { compressImage } from "@/utils/storage";

async function handleFile(file) {
  const compressedFile = await compressImage(file);
  // Upload compressedFile instead of original
}
```

## Storage Structure

Images are stored in Firebase Storage with this structure:

```
your-bucket.appspot.com/
└── products/
    ├── product_1704123456789_a1b2c3.jpg
    ├── product_1704123467890_d4e5f6.png
    └── product_1704123478901_g7h8i9.webp
```

**Filename format**: `product_{timestamp}_{random}.{ext}`

## Security Rules

Current rules (development):

```javascript
match /products/{imageId} {
  allow read: if true;
  allow write: if request.resource.size < 5 * 1024 * 1024
               && request.resource.contentType.matches('image/.*');
}
```

For production, add authentication:

```javascript
allow write: if request.auth != null
             && request.auth.token.admin == true
```

## Environment Variables

Required in `.env`:

```env
VITE_FIREBASE_API_KEY=your_api_key
VITE_FIREBASE_AUTH_DOMAIN=your-project.firebaseapp.com
VITE_FIREBASE_PROJECT_ID=your-project-id
VITE_FIREBASE_STORAGE_BUCKET=your-project.appspot.com
VITE_FIREBASE_MESSAGING_SENDER_ID=123456789012
VITE_FIREBASE_APP_ID=1:123456789012:web:abc123
```

## Troubleshooting

### Image won't upload

- Check Firebase config in `.env`
- Verify Storage is enabled in Firebase Console
- Check browser console for errors
- Ensure file is < 5MB and correct type

### "Firebase not initialized" error

- Restart dev server after updating `.env`
- Check that `firebase.js` is imported in `plugins/index.js`

### Storage rules error

- Go to Firebase Console > Storage > Rules
- Update rules to allow write access (see Security Rules above)

### Upload stuck at 0%

- Check internet connection
- Verify Storage bucket exists
- Check CORS configuration

## API Integration

The image URL from Firebase is stored in the product database:

```javascript
// After upload
const response = await fetch("/api/v1/products", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({
    name: "MacBook Pro",
    price: 1999,
    image_url: "https://firebasestorage.googleapis.com/...",
    // ... other fields
  }),
});
```

## Future Enhancements

- [ ] Multiple images per product
- [ ] Image cropping/editing
- [ ] Image gallery view
- [ ] Batch upload
- [ ] Watermarking
- [ ] CDN integration
- [ ] Lazy loading
- [ ] Thumbnail generation

## Benefits Over URL Strings

| Feature         | URL Strings     | Firebase Storage   |
| --------------- | --------------- | ------------------ |
| **Reliability** | Links can break | Permanent storage  |
| **Control**     | Depends on host | Full control       |
| **Speed**       | Variable        | Fast CDN delivery  |
| **Management**  | Manual tracking | Built-in tools     |
| **Security**    | Public only     | Configurable rules |
| **Scalability** | Limited         | Auto-scaling       |
| **Analytics**   | None            | Usage tracking     |

## Cost Estimation

Firebase Free Tier:

- **Storage**: 5GB (≈ 5,000 product images at 1MB each)
- **Downloads**: 1GB/day (≈ 1,000 image views/day at 1MB each)
- **Uploads**: 20,000/day

This is usually sufficient for small to medium e-commerce sites.

## Support

- [Firebase Documentation](https://firebase.google.com/docs/storage)
- [Firebase Console](https://console.firebase.google.com/)
- [Setup Guide](./FIREBASE_SETUP.md)
