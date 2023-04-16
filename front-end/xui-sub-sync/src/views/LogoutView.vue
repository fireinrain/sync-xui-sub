<template>
  <el-dialog
      v-model="dialogVisible"
      title="提示"
      width="30%"
  >
    <span>确定注销登录?</span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleLogout">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script>
import { defineComponent, ref } from 'vue';
import { ElButton, ElDialog } from 'element-plus';

const useLogoutDialog = () => {
  const that = this

  const dialogVisible = ref(true);

  const handleLogout = () => {
    // TODO: send request to logout API
    // after logout, navigate to the home page
    that.$router.push({ name: 'home' });
  };

  return {
    dialogVisible,
    handleLogout,
  };
};

export default defineComponent({
  name: 'LogoutPage',
  components: {
    ElButton,
    ElDialog,
  },
  setup() {
    const { dialogVisible, handleLogout } = useLogoutDialog();

    return { dialogVisible, handleLogout };
  },
});
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
