import Vue from 'vue';
import Vuex from 'vuex';

import apiClient from './api';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    alerts: {},
    artists: {},
    settings: {
      directories: {},
    },
  },
  mutations: {
    resetAlerts (state, alerts) {
      state.alerts = alerts;
    },
    resetArtists (state, artists) {
      state.artists = artists;
    },
    resetDirectories (state, directories) {
      state.settings.directories = directories;
    },
  },
  getters: {
    allAlerts: state => {
      return state.alerts;
    },
    allArtists: state => {
      return state.artists;
    },
    allDirectories: state => {
      return state.settings.directories;
    },
  },
  actions: {
    getArtists({commit}) {
      return new Promise((resolve) => {
        apiClient.getArtists().then((data) => {
          commit('resetArtists', data.items);
          resolve();
        });
      });
    },
    getDirectories({commit}) {
      return new Promise((resolve) => {
        apiClient.getDirectories().then((data) => {
          commit('resetDirectories', data.items);
          resolve();
        });
      });
    },
    pushDirectory({commit, state}, directory) {
      const directories = { ...state.settings.directories, [directory.id]: directory };

      return apiClient.pushDirectory(directory).then(() => {
        commit('resetDirectories', directories);
      });
    },
    removeDirectory({commit, state}, directory) {
      const directories = state.settings.directories.splice(
        state.settings.directories.indexOf(directory),
        1
      );

      return apiClient.removeDirectory(directory).then(() => {
        commit('resetDirectories', directories);
      });
    },
    triggerTrawl({commit}) {
      return new Promise((resolve) => {
        apiClient.triggerTrawl().then((data) => {
          commit('resetAlerts', {
            trawlFailure: (data.meta.errors > 0 && data.items.length < 1),
            trawlSuccess: (data.meta.errors == 0 && data.items.length > 0)
          });
          resolve();
        });
      });
    },
    updateArtist({commit, state}, artist) {
      const artists = { ...state.artists, [artist.text]: artist};

      return apiClient.updateArtist(artist).then(() => {
        commit('resetArtists', artists)
      });
    }
  }
});
