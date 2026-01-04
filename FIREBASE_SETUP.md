# Firebase Storage Setup Guide

This guide will help you set up Firebase Storage for product image uploads in your Jumia clone application.

## Step 1: Create a Firebase Project

1. Go to [Firebase Console](https://console.firebase.google.com/)
2. Click "Add project" or "Create a project"
3. Enter a project name (e.g., "jumia-clone")
4. Choose whether to enable Google Analytics (optional)
5. Click "Create project"

## Step 2: Register Your Web App

1. In your Firebase project, click on the Web icon (</>) to add a web app
2. Register your app with a nickname (e.g., "Jumia Clone Web")
3. Don't check "Firebase Hosting" (unless you want to use it)
4. Click "Register app"
5. Copy the Firebase configuration object

## Step 3: Enable Firebase Storage

1. In the Firebase Console, navigate to "Build" > "Storage"
2. Click "Get started"
3. Choose "Start in test mode" for development (you can secure it later)
4. Select a Cloud Storage location (choose the closest to your users)
5. Click "Done"

## Step 4: Configure Storage Security Rules (Important!)

For development, use these rules (in Firebase Console > Storage > Rules):

```javascript
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /products/{imageId} {
      // Allow anyone to read product images
      allow read: if true;

      // Allow authenticated users to write (you can add more restrictions)
      allow write: if request.resource.size < 5 * 1024 * 1024 // 5MB max
                   && request.resource.contentType.matches('image/.*');
    }
  }
}
```

For production, you should implement proper authentication:

```javascript
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /products/{imageId} {
      // Allow anyone to read product images
      allow read: if true;

      // Only allow admin users to upload (implement your own auth logic)
      allow write: if request.auth != null
                   && request.auth.token.admin == true
                   && request.resource.size < 5 * 1024 * 1024
                   && request.resource.contentType.matches('image/.*');
    }
  }
}
```

## Step 5: Update Environment Variables

1. Open the `.env` file in `console/v1/`
2. Replace the placeholder values with your actual Firebase config:

```env
VITE_FIREBASE_API_KEY=your_actual_api_key
VITE_FIREBASE_AUTH_DOMAIN=your-project-id.firebaseapp.com
VITE_FIREBASE_PROJECT_ID=your-project-id
VITE_FIREBASE_STORAGE_BUCKET=your-project-id.appspot.com
VITE_FIREBASE_MESSAGING_SENDER_ID=123456789012
VITE_FIREBASE_APP_ID=1:123456789012:web:abcdef1234567890
```

**Where to find these values:**

- Firebase Console > Project Settings > General > Your apps
- Scroll down to "SDK setup and configuration"
- Click "Config" radio button
- Copy the values from the config object

## Step 6: Restart Development Server

After updating the `.env` file:

```bash
cd console/v1
pnpm run dev
```

## How It Works

### Image Upload Flow

1. **User selects image**: File is validated for type (JPEG/PNG/WebP) and size (max 5MB)
2. **Client-side compression**: Image is compressed to reduce file size
3. **Preview shown**: User sees preview before uploading
4. **Upload to Firebase**: Image is uploaded with progress tracking
5. **Get download URL**: Firebase returns a permanent URL
6. **Save to database**: URL is saved with product data

### File Storage Structure

```
your-storage-bucket/
└── products/
    ├── product_1234567890_abc123.jpg
    ├── product_1234567891_def456.png
    └── product_1234567892_ghi789.webp
```

### Image Compression

- Maximum dimensions: 1200px (width or height)
- JPEG quality: 0.8 (80%)
- Automatic format preservation

## Security Best Practices

1. **File Size Validation**: Client validates 5MB max, server should also validate
2. **File Type Restriction**: Only allow image types (JPEG, PNG, WebP)
3. **Unique Filenames**: Timestamped filenames prevent overwriting
4. **Storage Rules**: Implement authentication in production
5. **CORS Configuration**: Already configured for web uploads

## Features Implemented

✅ **Image Upload Component**

- Drag and drop support
- File input with camera icon
- Real-time preview
- Upload progress bar
- File validation

✅ **Storage Utilities**

- `uploadProductImage()`: Upload with progress callback
- `deleteProductImage()`: Remove images from storage
- `compressImage()`: Client-side compression

✅ **Form Integration**

- Upload before submitting product
- Fallback to URL input if needed
- Preview uploaded images
- Remove and re-upload capability

## Testing the Upload

1. Navigate to `/admin/products/add`
2. Fill in product details
3. Click "Choose File" or drag an image
4. Click "Upload Image"
5. Watch the progress bar
6. Preview appears when done
7. Submit the form

## Troubleshooting

### Error: "Firebase not initialized"

- Check that `.env` has correct values
- Restart dev server after changing `.env`

### Error: "Permission denied"

- Check Firebase Storage Rules
- Ensure rules allow write access

### Error: "File too large"

- Maximum 5MB allowed
- Compress image before uploading

### Error: "Invalid file type"

- Only JPEG, PNG, WebP allowed
- Convert other formats first

### Upload stuck at 0%

- Check internet connection
- Verify Storage is enabled in Firebase Console
- Check browser console for errors

## Alternative: Using Manual URL

If you prefer not to use Firebase or need to test quickly:

1. Upload image to any image host (Imgur, Cloudinary, etc.)
2. Copy the direct image URL
3. Paste into the "Or enter image URL" field
4. Submit the form

## Next Steps

- [ ] Add image editing (crop, resize, filters)
- [ ] Support multiple images per product
- [ ] Implement image CDN for faster loading
- [ ] Add watermarking for product images
- [ ] Batch upload for multiple products

## Costs

Firebase offers a generous free tier:

- **Storage**: 5GB free
- **Downloads**: 1GB/day free
- **Uploads**: 20,000/day free

For most small to medium e-commerce sites, this is sufficient. Monitor usage in Firebase Console.

## Support

For issues:

- Check [Firebase Documentation](https://firebase.google.com/docs/storage)
- Review [Firebase Console](https://console.firebase.google.com/)
- Check browser console for errors
