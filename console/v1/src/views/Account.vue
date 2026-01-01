<template>
  <v-container class="pa-6 bg-grey-lighten-4" style="max-width: 1550px">
    <!-- Page Title -->
    <h1 class="text-h5 font-weight-regular mb-6">Account Overview</h1>

    <!-- Grid Layout for Cards -->
    <v-row>
      <!-- Account Details Card -->
      <v-col cols="12" md="6">
        <v-card elevation="0" class="pa-0">
          <v-card-title
            class="bg-grey-lighten-3 text-subtitle-1 font-weight-medium py-3"
          >
            ACCOUNT DETAILS
          </v-card-title>
          <v-card-text class="pa-6">
            <div class="text-body-1 font-weight-medium mb-1">victor thumbi</div>
            <div class="text-body-2">vickycrowny5@gmail.com</div>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Address Book Card -->
      <v-col cols="12" md="6">
        <v-card elevation="0" class="pa-0">
          <v-card-title
            class="bg-grey-lighten-3 text-subtitle-1 font-weight-medium py-3 d-flex justify-space-between align-center"
          >
            <span>ADDRESS BOOK</span>
            <v-icon color="orange" size="small">mdi-pencil</v-icon>
          </v-card-title>
          <v-card-text class="pa-6">
            <div class="text-body-2 font-weight-medium mb-2">
              Your default shipping address:
            </div>
            <div class="text-body-2">victor Waiguru</div>
            <div class="text-body-2">King'ong'o</div>
            <div class="text-body-2">Nyeri Town, Nyeri</div>
            <div class="text-body-2">+254 748218541 / +254 748218541</div>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Jumia Store Credit Card -->
      <v-col cols="12" md="6">
        <v-card elevation="0" class="pa-0">
          <v-card-title
            class="bg-grey-lighten-3 text-subtitle-1 font-weight-medium py-3"
          >
            JUMIA STORE CREDIT
          </v-card-title>
          <v-card-text class="pa-6">
            <div class="d-flex align-center">
              <v-icon color="blue" class="mr-2">mdi-credit-card-outline</v-icon>
              <span class="text-body-2">Jumia store credit balance: KSh 0</span>
            </div>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Newsletter Preferences Card -->
      <v-col cols="12" md="6">
        <v-card elevation="0" class="pa-0">
          <v-card-title
            class="bg-grey-lighten-3 text-subtitle-1 font-weight-medium py-3"
          >
            NEWSLETTER PREFERENCES
          </v-card-title>
          <v-card-text class="pa-6">
            <div class="text-body-2 mb-3">
              Manage your email communications to stay updated with the latest
              news and offers.
            </div>
            <a href="#" class="text-orange text-decoration-none text-body-2">
              Edit Newsletter preferences
            </a>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
// Account page
</script>
