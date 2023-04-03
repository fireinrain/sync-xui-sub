<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="submitForm">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username" required>

      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>

      <div>
        <img :src="captchaUrl" alt="captcha image">
        <input type="text" id="captcha" v-model="captcha" required>
      </div>

      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
      captcha: '',
      captchaUrl: ''
    }
  },
  methods: {
    async submitForm() {
      // validate captcha
      const response = await fetch('/api/captcha', {
        method: 'POST',
        body: JSON.stringify({captcha: this.captcha})
      });
      const result = await response.json();
      if (!result.success) {
        alert('Invalid captcha');
        return;
      }

      // login logic here
      const response2 = await fetch('/api/login', {
        method: 'POST',
        body: JSON.stringify({username: this.username, password: this.password})
      });
      const result2 = await response2.json();
      if (result2.success) {
        alert('Login successful');
      } else {
        alert('Invalid username or password');
      }
    },
    async getCaptcha() {
      const response = await fetch('/api/captcha');
      const result = await response.json();
      this.captchaUrl = result.url;
    }
  },
  mounted() {
    this.getCaptcha();
  }
}
</script>
