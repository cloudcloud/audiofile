<template>
  <div>
    <v-toolbar class="white--text" color="primary" app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-light">audiofile</span>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn flat to="/">
        <span class="mr-2 white--text">Artists</span>
      </v-btn>
      <v-btn flat to="/albums">
        <span class="mr-2 white--text">Albums</span>
      </v-btn>
      <v-btn flat to="/settings">
        <span class="mr-2 white--text">Settings</span>
      </v-btn>
      <v-btn flat @click="runProcess">
        <span class="mr-2 white--text">Re-process</span>
      </v-btn>

      <v-spacer></v-spacer>
      <v-btn flat href="https://github.com/cloudcloud/audiofile/releases/latest" target="_blank">
        <span class="mr-2 white--text">Latest Release</span>
      </v-btn>
    </v-toolbar>

    <v-container mt-5 pb-0 mb-2>
      <v-alert @click="trawlSuccess=false" v-model="trawlSuccess" type="success" transition="scale-transition">Successfully triggered the trawl!</v-alert>
      <v-alert @click="trawlFailure=false" v-model="trawlFailure" type="error" transition="scale-transition">Unable to trigger the trawl!</v-alert>
    </v-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      trawlFailure: false,
      trawlSuccess: false
    };
  },
  methods: {
    loadAlerts() {
      const alerts = this.$store.getters.allAlerts;
      this.trawlFailure = alerts.trawlFailure;
      this.trawlSuccess = alerts.trawlSuccess;
    },
    runProcess() {
      this.$store.dispatch('triggerTrawl').then(() => {
        this.loadAlerts();
      });
    }
  }
}
</script>
