<template>
  <v-sheet class="d-flex" style="background: #333; color: #fff">
    <v-container fluid class="pa-8">
      <v-row align="center" class="">
        <!-- Left: Logo -->
        <v-col cols="12" md="3" class="d-flex align-center">
          <div class="d-flex align-center">
            <div class="text-h4 font-weight-bold mr-3">JUMIA</div>
            <v-icon color="orange" size="32">mdi-star-circle</v-icon>
          </div>
        </v-col>

        <!-- Center: Newsletter form -->
        <v-col cols="12" md="6">
          <div class="d-flex flex-column">
            <div class="text-subtitle-1 font-weight-medium">NEW TO JUMIA?</div>
            <div class="caption mb-3">
              Subscribe to our newsletter to get updates on our latest offers!
            </div>

            <div class="d-flex align-center">
              <v-text-field
                v-model="email"
                placeholder="Enter E-mail Address"
                dense
                hide-details
                prepend-inner-icon="mdi-email"
                class="mr-3 bg-white"
                style="max-width: 420px"
                aria-label="email"
              />

              <v-btn variant='outlined' size='x-large' color="white" class="ml-2 pa-2" @click="subscribe"
                >Subscribe</v-btn
              >
            </div>

            <div class="d-flex align-center mt-3">
              <v-checkbox
                v-model="agreed"
                class="mr-2"
                hide-details
                density="compact"
                aria-label="agree"
              />
              <div class="caption">
                I agree to Jumia's Privacy and Cookie Policy. You can
                unsubscribe from newsletters at any time.
              </div>
            </div>

            <div class="caption mt-2" style="color: #fb8c00">
              I accept the Legal Terms
            </div>
          </div>
        </v-col>

        <!-- Right: Download app box -->
        <v-col cols="12" md="3" class="d-flex flex-column align-end">
          <div class="d-flex flex-column align-end">
            <div class="d-flex align-center mb-2">
              <v-avatar color="text"  size="36" class="mr-2 bg-primary pa-3">
                <v-icon dark>mdi-star</v-icon>
              </v-avatar>
              <div>
                <div class="subtitle-2 font-weight-medium">
                  DOWNLOAD JUMIA FREE APP
                </div>
                <div class="caption">Get access to exclusive offers!</div>
              </div>
            </div>

            <div class="d-flex ga-2">
              <v-btn
                
                variant=""
                class="d-flex align-center"
                large
                
              >
                <v-icon left color="text" size="x-large">mdi-apple</v-icon>
                App Store
              </v-btn>
              <v-btn
                
                variant=""
                class="d-flex align-center"
                small
              >
                <v-icon left color="text" size="x-large">mdi-google-play</v-icon>
                Google Play
              </v-btn>
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </v-sheet>
</template>

<script setup>
import { ref } from "vue";

const email = ref("");
const agreed = ref(false);

function subscribe() {
  // simple placeholder action
  if (!email.value) return;
  console.log("subscribe", { email: email.value, agreed: agreed.value });
  // Reset for now
  email.value = "";
  agreed.value = false;
}
</script>

<!-- No custom scoped CSS: using Vuetify inline utility classes -->
