<template>
  <div class="app-container" v-if="$route.path === '/project'">
    <Sidebar/>
    <div class="grp-container">
      <el-row>
        <h2 style="font-size: 16px; font-weight: bold; color: #132F51;">Project</h2>
      </el-row>

      <el-row style="width: 100%; display: flex; justify-content: flex-end;">
        <el-button type="primary" @click="form('create')">New Project</el-button>
      </el-row>

      <el-row></el-row>
    </div>
  </div>
  <router-view></router-view>
</template>
  
<script>
  import axios from 'axios';
  import Sidebar from '/src/components/Sidebar.vue';
  import { notification} from "/src/utils/notification";
  import Cookies from 'js-cookie';

  export default {
    name: 'Project',
    components: {
      Sidebar,
    },
    data() {
      return {
       
      };
    },
    mounted() {
    },
    methods: {
      form(action, data) {
        if (action == 'create') {
          this.$router.push({ path: '/project/form', query: { action: 'create' } });

          console.log(this.$router.getRoutes());
        }
      }
    }
  };
</script>

<style lang="scss" scoped>
  .app-container {
    display: flex;
    flex-direction: row;
  }

  .el-row {
    margin-bottom: 20px;
  }
  
  .el-row:last-child {
    margin-bottom: 0;
  }

  .content{
    .el-row {
      margin-bottom: 3px;
    }
  }

  label {
    font-size: 10px;
    color: #454545;
    font-weight: bold;
  }

  .font-12 {
    font-size: 11px;
  }

  .grp-center {
    width: 100px; height: 100px;
    color: #ffffff;
    border-radius: 100px;
    background-color: #132F51;
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>