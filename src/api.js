import axios from 'axios';

const client = axios.create({
  baseURL: 'http://localhost:8008',
  json: true,
});

const apiClient = {
  getArtists() {
    return this.perform('get', '/api/artists');
  },

  getArtist(artist) {
    return this.perform('get', `/api/artist/${artist}`);
  },

  getDirectories() {
    return this.perform('get', '/api/settings/directories');
  },

  pushDirectory(directory) {
    return this.perform('post', '/api/settings/directory', directory);
  },

  removeDirectory(directory) {
    return this.perform('delete', '/api/settings/directory', directory);
  },

  async perform(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
      headers: {
        "X-Client": 'Ahoy-hoy'
      }
    }).then(req => {
      return req.data;
    });
  }
};

export default apiClient;
