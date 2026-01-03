<template>
  <v-container class="pa-6 bg-grey-lighten-4" style="max-width: 1550px">
    <!-- Loading State -->
    <div v-if="authStore.loading" class="text-center py-12">
      <v-progress-circular indeterminate color="orange" size="64" />
      <p class="text-body-1 mt-4">Loading account details...</p>
    </div>

    <!-- Not Logged In -->
    <div v-else-if="!authStore.isAuthenticated" class="text-center py-12">
      <v-icon size="120" color="grey-lighten-1">mdi-account-alert</v-icon>
      <h2 class="text-h5 mt-4 mb-2">Please Log In</h2>
      <p class="text-body-1 text-grey-darken-1 mb-6">
        You need to be logged in to view your account details.
      </p>
      <v-btn color="orange" variant="flat" size="large" to="/login">
        Go to Login
      </v-btn>
    </div>

    <!-- Account Content -->
    <div v-else>
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
              <div class="text-body-1 font-weight-medium mb-1">
                {{ authStore.fullName }}
              </div>
              <div class="text-body-2">{{ authStore.user?.email }}</div>
              <div v-if="authStore.user?.phone" class="text-body-2 mt-2">
                {{ authStore.user.phone }}
              </div>
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
              <div class="text-body-2">{{ authStore.fullName }}</div>
              <div class="text-body-2">King'ong'o</div>
              <div class="text-body-2">Nyeri Town, Nyeri</div>
              <div class="text-body-2">
                {{ authStore.user?.phone || "No phone number" }}
              </div>
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
                <v-icon color="blue" class="mr-2"
                  >mdi-credit-card-outline</v-icon
                >
                <span class="text-body-2"
                  >Jumia store credit balance: KSh 0</span
                >
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
    </div>
  </v-container>
</template>

<script setup>
import { onMounted } from "vue";
import { useAuthStore } from "@/stores/auth";

const authStore = useAuthStore();

onMounted(async () => {
  // Fetch user data if authenticated but user data is missing
  if (authStore.isAuthenticated && !authStore.user) {
    await authStore.fetchUser();
  }
});
</script>
