<template>
  <div>
    <h1 class="text-h5 font-weight-regular mb-6">Newsletter Preferences</h1>
    <v-card elevation="0">
      <v-card-text class="pa-6">
        <p class="text-body-1 mb-4">
          Manage your email communications to stay updated with the latest news
          and offers.
        </p>
        <v-divider class="my-4"></v-divider>
        <v-checkbox
          label="Receive promotional emails"
          color="orange"
          hide-details
          class="mb-2"
        ></v-checkbox>
        <v-checkbox
          label="Receive order updates"
          color="orange"
          hide-details
          class="mb-2"
        ></v-checkbox>
        <v-checkbox
          label="Receive newsletters"
          color="orange"
          hide-details
          class="mb-4"
        ></v-checkbox>
        <v-btn color="orange" variant="flat" class="mt-2">
          Save Preferences
        </v-btn>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup>
// Newsletter Preferences page
</script>
