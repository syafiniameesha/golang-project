<template>
    <el-menu
      :default-active="currentRoute"
      class="el-menu-vertical-demo side-menu"
      @select="handleMenuClick"
    >
        <el-menu-item index="/homepage">Home</el-menu-item>
        <el-menu-item index="/userManagement">User Management</el-menu-item>
        <el-menu-item index="/logout" @click="logout">Logout</el-menu-item>
    </el-menu>
  </template>
  
<script>
import { useRouter, useRoute } from 'vue-router';
import { ElMenu, ElMenuItem } from 'element-plus';
import { notification} from "/src/utils/notification";
import Cookies from 'js-cookie';
export default {
    name: 'Sidebar',
    components: {
        ElMenu,
        ElMenuItem,
    },
    data() {
        return {
            currentRoute: '',
        };
    },
    mounted() {
        this.currentRoute = this.$route.path;
    },
    methods: {
        handleMenuClick(index) {
            if (index !== '/logout') {
                this.$router.push(index);
            }
        },
        logout() {
            notification('Logging Out', 'You will logout from the system', 'bottom-left');

            Cookies.remove('token');

            setTimeout(() => {
                this.$router.push('/');
                location.reload();
            }, 1000);
        },
    },
};
</script>
  
<style lang="scss" scoped>
.side-menu {
    width: 160px;
    padding-top: 20px;
}

.side-menu .el-menu-item.is-active {
    background-color: #132F51 !important;
    color: #fff !important;
}
    
</style>
  