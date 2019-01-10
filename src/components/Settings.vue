<template>
  <div>
    <v-container fluid grid-list-sm>
      <v-layout row wrap>
        <v-flex xs12 md-5>
          <h2 class="display-1 font-weight-bold mb-3">settings</h2>
        </v-flex>
      </v-layout>
    </v-container>

    <v-toolbar flat color="white">
      <v-toolbar-title>Directories</v-toolbar-title>
      <v-divider class="mx-2" inset vertical></v-divider>
      <v-spacer></v-spacer>

      <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" dark class="mb-2">Add Directory</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="editedItem.directory" label="Directory"></v-text-field>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
            <v-btn color="blue darken-1" flat @click="save">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-toolbar>
    <v-data-table :headers="head" :items="directories" class="elevation-1">
      <template slot="items" slot-scope="props">
        <td>{{ props.item.directory }}</td>
        <td class="justify-center layout px-0">
          <v-icon small class="mr-2" @click="editItem(props.item)">edit</v-icon>
          <v-icon small @click="deleteItem(props.item)">delete</v-icon>
        </td>
      </template>
      <template slot="no-data">
        <v-btn color="primary" @click="loadDirectories">Reset</v-btn>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    defaultItem: {id: '', directory: ''},
    dialog: false,
    directories: [],
    editedIndex: -1,
    editedItem: {id: '', directory: ''},
    head: [{text: 'Directory', align: 'left', value: 'directory'}, {text: 'Actions', sortable: false, value: 'directory'}]
  }),
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Directory': 'Edit Directory';
    },
    ...mapGetters(['allDirectories']),
  },
  watch: {
    dialog(val) {
      val || this.close();
    }
  },
  created() {
    this.$store.dispatch('getDirectories').then(() => {
      this.loadDirectories();
    });
  },
  methods: {
    loadDirectories() {
      this.directories = this.$store.getters.allDirectories;
    },
    editItem(item) {
      this.editedIndex = this.directories.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },
    deleteItem(item) {
      if (confirm('Are you sure you want to delete this directory?')) {
        this.$store.dispatch('removeDirectory', item).then(() => {
          // display a confirmation dialog
        });
      }
    },
    close() {
      this.dialog = false;
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      }, 300)
    },
    save() {
      this.$store.dispatch('pushDirectory', this.editedItem).then(() => {
        // display confirmation dialog
      });

      if (this.editedIndex > -1) {
        Object.assign(this.directories[this.editedIndex], this.editedItem);
      } else {
        this.directories.push(this.editedItem);
      }
      this.close();
    },
    ...mapMutations(['resetDirectories']),
    ...mapActions(['getDirectories']),
  }
};
</script>

<style>

</style>

