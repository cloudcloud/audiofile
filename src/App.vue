<template>
  <v-app>
    <Header />

    <router-view />

    <Footer />
  </v-app>
</template>

<script>
import Footer from './components/Footer'
import Header from './components/Header'

export default {
  name: 'App',
  components: {
    Header,
    Footer
  },
  data () {
    return {
      //
    }
  }
}
</script>
