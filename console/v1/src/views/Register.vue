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
            Join Jumia Today
          </h1>
          <p class="text-h6 text-white">Create an account and start shopping</p>
        </div>
      </v-col>

      <!-- Right Side - Register Form -->
      <v-col cols="12" md="6" class="d-flex align-center justify-center py-8">
        <v-card
          elevation="8"
          max-width="500"
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
          <div class="text-center mb-6">
            <h2 class="text-h5 font-weight-bold mb-2">Create Account</h2>
            <p class="text-body-2 text-grey-darken-1">Sign up to get started</p>
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

          <!-- Success Alert -->
          <v-alert
            v-if="successMessage"
            type="success"
            variant="tonal"
            closable
            class="mb-4"
            @click:close="successMessage = ''"
          >
            {{ successMessage }}
          </v-alert>

          <!-- Register Form -->
          <v-form ref="form" @submit.prevent="handleRegister">
            <v-row>
              <v-col cols="12" sm="6">
                <v-text-field
                  v-model="formData.first_name"
                  label="First Name"
                  variant="outlined"
                  prepend-inner-icon="mdi-account-outline"
                  :rules="nameRules"
                  :disabled="authStore.loading"
                  required
                />
              </v-col>
              <v-col cols="12" sm="6">
                <v-text-field
                  v-model="formData.last_name"
                  label="Last Name"
                  variant="outlined"
                  :rules="nameRules"
                  :disabled="authStore.loading"
                  required
                />
              </v-col>
            </v-row>

            <v-text-field
              v-model="formData.email"
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
              v-model="formData.phone"
              label="Phone Number (Optional)"
              variant="outlined"
              prepend-inner-icon="mdi-phone-outline"
              :disabled="authStore.loading"
              class="mb-3"
              placeholder="+254 700 000 000"
            />

            <v-text-field
              v-model="formData.password"
              label="Password"
              :type="showPassword ? 'text' : 'password'"
              variant="outlined"
              prepend-inner-icon="mdi-lock-outline"
              :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
              :rules="passwordRules"
              :disabled="authStore.loading"
              required
              class="mb-3"
              @click:append-inner="showPassword = !showPassword"
            />

            <v-text-field
              v-model="confirmPassword"
              label="Confirm Password"
              :type="showPassword ? 'text' : 'password'"
              variant="outlined"
              prepend-inner-icon="mdi-lock-outline"
              :rules="confirmPasswordRules"
              :disabled="authStore.loading"
              required
              class="mb-4"
            />

            <!-- Terms & Conditions -->
            <v-checkbox
              v-model="acceptTerms"
              color="orange"
              :rules="[
                (v) => !!v || 'You must accept the terms and conditions',
              ]"
              :disabled="authStore.loading"
              class="mb-2"
            >
              <template #label>
                <span class="text-body-2">
                  I agree to the
                  <a href="#" class="text-orange" @click.prevent
                    >Terms & Conditions</a
                  >
                </span>
              </template>
            </v-checkbox>

            <!-- Register Button -->
            <v-btn
              type="submit"
              block
              size="large"
              color="orange"
              variant="flat"
              :loading="authStore.loading"
              class="mb-4"
            >
              Create Account
            </v-btn>

            <!-- Divider -->
            <v-divider class="my-4" />

            <!-- Login Link -->
            <div class="text-center">
              <span class="text-body-2 text-grey-darken-1">
                Already have an account?
              </span>
              <v-btn
                variant="text"
                color="orange"
                size="small"
                class="ml-1"
                @click="$router.push('/login')"
              >
                Sign In
              </v-btn>
            </div>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const router = useRouter();
const authStore = useAuthStore();

const form = ref(null);
const showPassword = ref(false);
const acceptTerms = ref(false);
const confirmPassword = ref("");
const successMessage = ref("");

const formData = reactive({
  first_name: "",
  last_name: "",
  email: "",
  phone: "",
  password: "",
});

const nameRules = [
  (v) => !!v || "Name is required",
  (v) => v.length >= 2 || "Name must be at least 2 characters",
];

const emailRules = [
  (v) => !!v || "Email is required",
  (v) => /.+@.+\..+/.test(v) || "Email must be valid",
];

const passwordRules = [
  (v) => !!v || "Password is required",
  (v) => v.length >= 6 || "Password must be at least 6 characters",
];

const confirmPasswordRules = [
  (v) => !!v || "Please confirm your password",
  (v) => v === formData.password || "Passwords do not match",
];

async function handleRegister() {
  const { valid } = await form.value.validate();

  if (!valid) {
    return;
  }

  try {
    await authStore.register(formData);

    successMessage.value = "Account created successfully! Redirecting...";

    // Redirect to home after short delay
    setTimeout(() => {
      router.push("/");
    }, 2000);
  } catch (err) {
    // Error is handled in the store and displayed in the alert
    console.error("Registration failed:", err);
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
