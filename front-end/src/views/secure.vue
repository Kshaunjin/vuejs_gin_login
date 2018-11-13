<template>
  <div id="secure">
     <h1>Secure Area</h1>
        <p>
           This is a secure area
        </p>
      <div id="app2">
       <p>{{ datas }}</p>
       <button @click="getData">取得</button>
      </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'Secure',
  data () {
    return {
      datas: [ ]
    }
  },

  methods: {
    getData: function () {
      let self = this
      axios({
        methods: 'get',
        url: 'http://localhost:8000/api/info',
        headers: {
          Authorization: this.$cookies.get('token')
        }
      })
        .then((resp) => {
          self.datas = resp.data
        })
        .catch(error => {
          console.log(error.response)
        })
    }
  }

}
</script>
<style scoped>
    #secure {
        background-color: #FFFFFF;
        border: 1px solid #CCCCCC;
        padding: 20px;
        margin-top: 10px;
    }
</style>
