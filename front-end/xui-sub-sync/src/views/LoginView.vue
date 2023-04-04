<template>
  <div class="login-wrapper">
    <el-form ref="loginForm" :model="loginForm" label-width="80px" class="login-form">
      <h3 class="login-title">xui-sub-sync</h3>
      <el-form-item label="Username" prop="username">
        <el-input v-model="loginForm.username" placeholder="Username"></el-input>
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input v-model="loginForm.password" placeholder="Password" type="password"></el-input>
      </el-form-item>
      <el-form-item label="Image Recaptcha" prop="recaptcha">
        <!-- replace the placeholder with your image recaptcha component -->
        <div style="height: 100px;">Image Recaptcha Component Placeholder</div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleSubmit" class="login-button">Login</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { ref } from 'vue';
import { ElForm, ElFormItem, ElInput, ElButton } from 'element-plus';

export default {
  name: 'LoginView',
  components: {
    ElForm,
    ElFormItem,
    ElInput,
    ElButton,
  },
  setup() {
    const loginForm = ref({
      username: '',
      password: '',
      recaptcha: '',
    });

    const handleSubmit = () => {
      // handle form submission here
    };

    return {
      loginForm,
      handleSubmit,
    };
  },
};
</script>

<style scoped>
.login-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.login-form {
  width: 400px;
}

.login-title {
  text-align: center;
  margin-bottom: 20px;
  font-size: 24px;
}
/* Set the width of the login button to be the same as the username field */
.el-form-item__content .el-input__inner,
.login-button {
  width: 100%;
}
</style>
