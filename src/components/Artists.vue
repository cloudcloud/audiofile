<template>
  <v-container fluid grid-list-sm>
    <v-layout row mt-5 wrap>
      <v-flex mb-3>
        <h2 class="display-1 font-weight-bold mb-3">artists</h2>
      </v-flex>

      <v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="artists" class="elevation-1">
            <template slot="items" slot-scope="props">
              <td><router-link :to="props.item.href" :href="props.item.href">{{props.item.text}}</router-link></td>
              <td>{{ props.item.albums.length }}</td>
              <td>{{ props.item.status }}</td>
            </template>
          </v-data-table>
        </v-flex>
      </v-container>

    </v-layout>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data() {
    return  {
      headers: [
        {text: 'Artist', align: 'left', value: 'text'},
        {text: 'Albums', value: 'albums'},
        {text: 'Status', value: 'status'}
      ],
      artists: [],
    }
  },
  created() {
    this.$store.dispatch('getArtists').then(() => {
      this.loadArtists();
    });
  },
  methods: {
    loadArtists() {
      this.artists = this.$store.getters.allArtists;
    },
    ...mapMutations(['resetArtists']),
    ...mapActions(['getArtists']),
  },
  computed: {
    ...mapGetters(['allArtists']),
  },
};
</script>

<style>

</style>
