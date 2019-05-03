<template>
  <div class="home">
    <img
      alt="Vue logo"
      src="../assets/logo.png"
    >
    <HelloWorld :msg="this.$store.state.resp" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import axios from 'axios'

export default {
  name: 'home',
  components: {
    HelloWorld
  },
  mounted: function () {
    // var code = this.$route.query.code
    axios.get('http://127.0.0.1:3030/hello')
      .then((resp) => {
        this.$store.dispatch('response', resp.data)
      })
  }
}
</script>
