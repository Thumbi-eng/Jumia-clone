<template>
  <div>
    <h1 class="text-h5 font-weight-regular mb-6">Close Account</h1>
    <v-card elevation="0">
      <v-card-text class="pa-6">
        <v-alert type="warning" variant="tonal" class="mb-4">
          Warning: Closing your account is permanent and cannot be undone.
        </v-alert>
        <p class="text-body-1 mb-4">
          Are you sure you want to close your Jumia account? This action will:
        </p>
        <ul class="mb-4">
          <li>Delete all your personal information</li>
          <li>Remove your order history</li>
          <li>Cancel any active vouchers</li>
          <li>Remove you from our mailing list</li>
        </ul>
        <v-btn color="error" variant="flat">Close My Account</v-btn>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup>
// Close Account page
</script>
