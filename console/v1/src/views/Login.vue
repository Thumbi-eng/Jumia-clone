<template>
  <v-container fluid class="fill-height bg-grey-lighten-3 pa-0">
    <v-row no-gutters class="fill-height">
      <!-- Left Side - Image/Brand -->
      <v-col
        cols="12"
        md="6"
        class="d-none d-md-flex align-center justify-center bg-orange"
      >
        <div class="text-center pa-8">
          <v-img
            src="/Jumia-Logo.jpg"
            alt="Jumia Logo"
            max-width="300"
            class="mx-auto mb-8"
          />
          <h1 class="text-h3 text-white font-weight-bold mb-4">
            Welcome to Jumia
          </h1>
          <p class="text-h6 text-white">
            Shop millions of products at the best prices
          </p>
        </div>
      </v-col>

      <!-- Right Side - Login Form -->
      <v-col cols="12" md="6" class="d-flex align-center justify-center">
        <v-card
          elevation="8"
          max-width="450"
          width="100%"
          class="pa-8"
          rounded="lg"
        >
          <!-- Mobile Logo -->
          <div class="d-md-none text-center mb-6">
            <v-img
              src="/Jumia-Logo.jpg"
              alt="Jumia Logo"
              max-width="150"
              class="mx-auto"
            />
          </div>

          <!-- Welcome Text -->
          <div class="text-center mb-8">
            <h2 class="text-h5 font-weight-bold mb-2">Welcome Back</h2>
            <p class="text-body-2 text-grey-darken-1">
              Sign in to continue shopping
            </p>
          </div>

          <!-- Error Alert -->
          <v-alert
            v-if="authStore.error"
            type="error"
            variant="tonal"
            closable
            class="mb-4"
            @click:close="authStore.clearError"
          >
            {{ authStore.error }}
          </v-alert>

          <!-- Login Form -->
          <v-form ref="form" @submit.prevent="handleLogin">
            <v-text-field
              v-model="email"
              label="Email Address"
              type="email"
              variant="outlined"
              prepend-inner-icon="mdi-email-outline"
              :rules="emailRules"
              :disabled="authStore.loading"
              required
              class="mb-3"
            />

            <v-text-field
              v-model="password"
              label="Password"
              :type="showPassword ? 'text' : 'password'"
              variant="outlined"
              prepend-inner-icon="mdi-lock-outline"
              :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
              :rules="passwordRules"
              :disabled="authStore.loading"
              required
              class="mb-2"
              @click:append-inner="showPassword = !showPassword"
            />

            <!-- Remember Me & Forgot Password -->
            <div class="d-flex justify-space-between align-center mb-6">
              <v-checkbox
                v-model="rememberMe"
                label="Remember me"
                hide-details
                density="compact"
                color="orange"
              />
              <v-btn
                variant="text"
                color="orange"
                size="small"
                @click="$router.push('/forgot-password')"
              >
                Forgot Password?
              </v-btn>
            </div>

            <!-- Login Button -->
            <v-btn
              type="submit"
              block
              size="large"
              color="orange"
              variant="flat"
              :loading="authStore.loading"
              class="mb-4"
            >
              Sign In
            </v-btn>

            <!-- Divider -->
            <v-divider class="my-6" />

            <!-- Register Link -->
            <div class="text-center">
              <span class="text-body-2 text-grey-darken-1">
                Don't have an account?
              </span>
              <v-btn
                variant="text"
                color="orange"
                size="small"
                class="ml-1"
                @click="$router.push('/register')"
              >
                Sign Up
              </v-btn>
            </div>
          </v-form>

          <!-- Demo Credentials -->
          <v-alert type="info" variant="tonal" density="compact" class="mt-6">
            <div class="text-caption">
              <strong>Demo Account:</strong><br />
              Email: test@example.com<br />
              Password: password123
            </div>
          </v-alert>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const router = useRouter();
const authStore = useAuthStore();

const form = ref(null);
const email = ref("");
const password = ref("");
const showPassword = ref(false);
const rememberMe = ref(false);

const emailRules = [
  (v) => !!v || "Email is required",
  (v) => /.+@.+\..+/.test(v) || "Email must be valid",
];

const passwordRules = [
  (v) => !!v || "Password is required",
  (v) => v.length >= 6 || "Password must be at least 6 characters",
];

async function handleLogin() {
  const { valid } = await form.value.validate();

  if (!valid) {
    return;
  }

  try {
    await authStore.login(email.value, password.value);

    // Redirect to home or intended page
    const redirectTo = router.currentRoute.value.query.redirect || "/";
    router.push(redirectTo);
  } catch (err) {
    // Error is handled in the store and displayed in the alert
    console.error("Login failed:", err);
  }
}
</script>

<style scoped>
.fill-height {
  min-height: 100vh;
}

.bg-orange {
  background: linear-gradient(135deg, #ff6b35 0%, #f7931e 100%);
}
</style>
