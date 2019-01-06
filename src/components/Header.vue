<template>
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

    <v-spacer></v-spacer>
    <v-btn flat href="https://github.com/cloudcloud/audiofile/releases/latest" target="_blank">
      <span class="mr-2 white--text">Latest Release</span>
    </v-btn>
  </v-toolbar>
</template>
