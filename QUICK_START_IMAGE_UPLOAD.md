# Quick Start: Testing Image Upload

## Prerequisites Checklist

- [x] Firebase SDK installed (`firebase` package)
- [x] Storage utilities created (`src/utils/storage.js`)
- [x] Firebase plugin created (`src/plugins/firebase.js`)
- [x] AddProducts.vue updated with upload component
- [x] .gitignore updated to exclude `.env`

## Setup Steps

### 1. Configure Firebase (5 minutes)

**Option A: Create New Firebase Project**

1. Go to https://console.firebase.google.com/
2. Click "Add project"
3. Name it "jumia-clone" (or any name)
4. Disable Google Analytics (optional)
5. Click "Create project"

**Option B: Use Existing Project**

Skip to step 2 if you already have a Firebase project.

### 2. Enable Storage (2 minutes)

1. In Firebase Console, click "Storage" in left sidebar
2. Click "Get started"
3. Select "Start in test mode" (we'll secure it later)
4. Choose a location close to you
5. Click "Done"

### 3. Get Firebase Configuration (1 minute)

1. Click the gear icon ⚙️ next to "Project Overview"
2. Select "Project settings"
3. Scroll to "Your apps" section
4. If no web app exists:
   - Click the web icon `</>`
   - Name it "Jumia Clone Web"
   - Click "Register app"
5. Copy the `firebaseConfig` object values

### 4. Update Environment Variables (1 minute)

Edit `console/v1/.env` and replace with your actual values:

```env
export BASE_URL=http://localhost:3000/
VITE_FIREBASE_API_KEY=AIzaSyXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
VITE_FIREBASE_AUTH_DOMAIN=your-project-id.firebaseapp.com
VITE_FIREBASE_PROJECT_ID=your-project-id
VITE_FIREBASE_STORAGE_BUCKET=your-project-id.appspot.com
VITE_FIREBASE_MESSAGING_SENDER_ID=123456789012
VITE_FIREBASE_APP_ID=1:123456789012:web:abcdef123456
```

### 5. Restart Dev Server

```bash
cd console/v1

# Kill existing server (if running)
pkill -f "vite"

# Start fresh
pnpm run dev
```

## Testing the Feature

### Test 1: Upload Image File

1. Open browser to http://localhost:5173
2. Navigate to **Admin > Add Products** or go to `/admin/products/add`
3. Fill in basic product details:

   - Name: "Test Product with Upload"
   - Description: "Testing Firebase image upload"
   - Price: 99
   - Category: "Electronics"
   - Stock: 10
   - Brand: "TestBrand"

4. **Upload Image**:

   - Click "Choose File" button
   - Select a JPEG/PNG/WebP image (< 5MB)
   - Preview should appear immediately
   - Click "Upload Image" button
   - Watch progress bar (should reach 100%)
   - Success message appears

5. Click "Add Product"
6. Verify product was created

### Test 2: Drag & Drop

1. Repeat steps 1-3 from Test 1
2. Instead of clicking "Choose File":
   - Open file explorer
   - Drag an image file
   - Drop it on the upload area
3. Continue with upload and submit

### Test 3: Manual URL (Fallback)

1. Repeat steps 1-3 from Test 1
2. Skip file upload
3. Paste an image URL in "Or enter image URL" field:
   ```
   https://images.unsplash.com/photo-1505740420928-5e560c06d30e
   ```
4. Click "Add Product"
5. Verify product was created

### Test 4: File Validation

**Too Large File**:

1. Try to upload a file > 5MB
2. Should show error: "Image size must be less than 5MB"

**Wrong Type**:

1. Try to upload a PDF or other non-image file
2. Should show error: "Please select a valid image file"

### Test 5: Remove Image

1. Upload an image (see Test 1)
2. Click "Remove Image" button
3. Preview should disappear
4. Can upload a different image

## Verify in Firebase Console

1. Go to Firebase Console > Storage
2. Click on the `products/` folder
3. You should see uploaded images:
   - `product_1704123456_abc123.jpg`
   - `product_1704123789_def456.png`
4. Click an image to see details
5. Copy "Download URL" and paste in browser
6. Image should load

## Verify in Database

Check that the Firebase URL was saved:

```bash
# In backend directory
cd jumia-clone-backend/services/product-service

# Connect to PostgreSQL
docker exec -it jumia_postgres psql -U postgres -d jumia_products

# Query products
SELECT id, name, image_url FROM products ORDER BY created_at DESC LIMIT 5;
```

You should see Firebase URLs like:

```
https://firebasestorage.googleapis.com/v0/b/your-project.appspot.com/o/products%2Fproduct_1234567890_abc123.jpg?alt=media&token=...
```

## Troubleshooting

### Error: "Firebase not initialized"

**Cause**: Environment variables not loaded or dev server not restarted

**Fix**:

```bash
# Verify .env file exists
cat console/v1/.env

# Kill all node processes
pkill -f "vite"
pkill -f "node"

# Restart
cd console/v1
pnpm run dev
```

### Error: "Permission denied" during upload

**Cause**: Storage rules too restrictive

**Fix**:

1. Go to Firebase Console > Storage > Rules
2. Replace with:

```javascript
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /products/{imageId} {
      allow read: if true;
      allow write: if request.resource.size < 5 * 1024 * 1024
                   && request.resource.contentType.matches('image/.*');
    }
  }
}
```

3. Click "Publish"

### Upload stuck at 0%

**Cause**: Network issue or invalid Firebase config

**Fix**:

1. Check internet connection
2. Verify Firebase config in `.env`
3. Check browser console for errors
4. Try uploading a smaller image

### Image URL not saving to database

**Cause**: Backend API not receiving correct data

**Fix**:

1. Open browser DevTools > Network tab
2. Submit product form
3. Check the request payload
4. Verify `image_url` field has the Firebase URL
5. Check backend logs for errors

## Browser Console Debugging

Press F12 to open DevTools, check for:

**Success**:

```
Image uploaded successfully!
Download URL: https://firebasestorage.googleapis.com/...
```

**Errors**:

```
Error uploading image: [error message]
Firebase initialization failed: [error message]
```

## Quick Commands Reference

```bash
# Install dependencies
cd console/v1
pnpm install

# Run dev server
pnpm run dev

# Check if dev server is running
lsof -i :5173

# Kill dev server
pkill -f "vite"

# Check Firebase package
pnpm list firebase

# View environment variables
cat console/v1/.env
```

## Success Indicators

✅ **Upload Working** when you see:

- Image preview appears after selection
- Progress bar moves from 0% to 100%
- Success message shows "Image uploaded successfully!"
- Firebase URL appears in form data
- Product created with image URL in database
- Image visible in Firebase Console > Storage

## Next Steps After Testing

1. ✅ Upload test images successfully
2. ✅ Verify images in Firebase Storage
3. ✅ Verify URLs in database
4. ✅ Check images display on product pages
5. 🔄 Update Storage Rules for production
6. 🔄 Add authentication for admin routes
7. 🔄 Implement image deletion when products are deleted
8. 🔄 Add image editing features (crop, resize)

## Performance Monitoring

After several uploads, check Firebase Console:

- Storage usage (should be < 5GB on free tier)
- Number of files uploaded
- Download bandwidth used

## Need Help?

- Check [FIREBASE_SETUP.md](./FIREBASE_SETUP.md) for detailed setup
- Check [IMAGE_UPLOAD_FEATURE.md](./IMAGE_UPLOAD_FEATURE.md) for feature docs
- Review browser console for specific errors
- Check Firebase Console for storage issues
