<template>
    <div class="login">
        <label>
            Username
        </label>
        <input type="text" id="username" v-model="username">
        <label>
            Password
        </label>
        <input type="password" id="password" v-model="password">
        <button v-on:click="login">
            Log in
        </button>
    </div>
</template>

<script>
 export default {
     name: 'Login',
     data: function() {
         return {
             username: '',
             password: ''
         }
     },
     methods: {
         login: function() {
             this.axios.post("http://localhost:8081/users/login", {
                 username: this.username,
                 password: this.password
             }).then(function(response) {
                 response
             })
         }
     }
 }
</script>

<style>
</style>
