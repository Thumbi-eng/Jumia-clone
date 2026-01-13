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
                @click="forgotPasswordDialog = true"
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
              class="mb-3"
            >
              Sign In
            </v-btn>

            <!-- Demo Login Button -->
            <v-btn
              variant="outlined"
              color="orange"
              size="large"
              block
              :loading="demoLoading"
              @click="handleDemoLogin"
              class="mb-4"
            >
              <v-icon class="mr-2">mdi-test-tube</v-icon>
              Demo Login (No Backend Required)
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

    <!-- Forgot Password Dialog -->
    <v-dialog v-model="forgotPasswordDialog" max-width="500" persistent>
      <v-card>
        <v-card-title class="text-h5 font-weight-bold pa-6">
          Reset Password
        </v-card-title>

        <v-card-text class="px-6">
          <v-alert
            v-if="resetSuccess"
            type="success"
            variant="tonal"
            class="mb-4"
          >
            Password reset email sent! Check your inbox.
          </v-alert>

          <v-alert
            v-if="resetError"
            type="error"
            variant="tonal"
            class="mb-4"
            closable
            @click:close="resetError = null"
          >
            {{ resetError }}
          </v-alert>

          <p class="text-body-2 text-grey-darken-1 mb-4">
            Enter your email address and we'll send you a link to reset your
            password.
          </p>

          <v-form ref="resetForm" @submit.prevent="handleResetPassword">
            <v-text-field
              v-model="resetEmail"
              label="Email Address"
              type="email"
              variant="outlined"
              prepend-inner-icon="mdi-email-outline"
              :rules="emailRules"
              :disabled="resetLoading || resetSuccess"
              required
              autofocus
            />
          </v-form>
        </v-card-text>

        <v-card-actions class="px-6 pb-6">
          <v-spacer />
          <v-btn
            variant="text"
            @click="closeForgotPasswordDialog"
            :disabled="resetLoading"
          >
            Cancel
          </v-btn>
          <v-btn
            color="orange"
            variant="flat"
            :loading="resetLoading"
            :disabled="resetSuccess"
            @click="handleResetPassword"
          >
            {{ resetSuccess ? "Email Sent" : "Send Reset Link" }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { getAuth, sendPasswordResetEmail } from "firebase/auth";

const router = useRouter();
const authStore = useAuthStore();

const form = ref(null);
const email = ref("");
const password = ref("");
const showPassword = ref(false);
const rememberMe = ref(false);
const demoLoading = ref(false);

// Forgot password
const forgotPasswordDialog = ref(false);
const resetForm = ref(null);
const resetEmail = ref("");
const resetLoading = ref(false);
const resetSuccess = ref(false);
const resetError = ref(null);

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

async function handleDemoLogin() {
  demoLoading.value = true;

  try {
    // Simulate demo login without backend
    const demoUser = {
      id: "demo-user-123",
      email: "demo@jumia.com",
      first_name: "Demo",
      last_name: "User",
      phone: "+254712345678",
      role: "customer",
      is_admin: false,
      created_at: new Date().toISOString(),
    };

    const demoToken = "demo-token-" + Date.now();

    // Manually set auth state
    authStore.user = demoUser;
    authStore.accessToken = demoToken;

    // Store in localStorage
    localStorage.setItem("user", JSON.stringify(demoUser));
    localStorage.setItem("access_token", demoToken);

    // Clear any errors
    authStore.clearError();

    // Redirect to home
    setTimeout(() => {
      router.push("/");
    }, 500);
  } catch (err) {
    console.error("Demo login failed:", err);
  } finally {
    demoLoading.value = false;
  }
}

async function handleResetPassword() {
  const { valid } = await resetForm.value.validate();

  if (!valid) {
    return;
  }

  resetLoading.value = true;
  resetError.value = null;
  resetSuccess.value = false;

  try {
    const auth = getAuth();
    await sendPasswordResetEmail(auth, resetEmail.value);
    resetSuccess.value = true;

    // Close dialog after 3 seconds
    setTimeout(() => {
      closeForgotPasswordDialog();
    }, 3000);
  } catch (err) {
    console.error("Password reset error:", err);

    // Handle Firebase errors
    switch (err.code) {
      case "auth/user-not-found":
        resetError.value = "No account found with this email address.";
        break;
      case "auth/invalid-email":
        resetError.value = "Invalid email address.";
        break;
      case "auth/too-many-requests":
        resetError.value = "Too many requests. Please try again later.";
        break;
      default:
        resetError.value = err.message || "Failed to send reset email.";
    }
  } finally {
    resetLoading.value = false;
  }
}

function closeForgotPasswordDialog() {
  forgotPasswordDialog.value = false;
  resetEmail.value = "";
  resetSuccess.value = false;
  resetError.value = null;
  if (resetForm.value) {
    resetForm.value.resetValidation();
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
