<template>
    <div class="signup">
        <label>
            Username
        </label>
        <input type="text" id="username" v-model="username">
        <label>
            Password
        </label>
        <input type="password" id="password" v-model="password">
        <button v-on:click="signup">
            Sign up
        </button>
    </div>
</template>

<script>
 export default {
     name: 'Signup',
     data: function() {
         return {
             username: '',
             password: ''
         }
     },
     methods: {
         login: function() {
             this.axios.post("http://localhost:8081/users", {
                 username: this.username,
                 password: this.password
             }).then(function(response) {
                response
                this.$route.go("/login");
             })
         }
     }
 }
</script>

<style>
</style>
