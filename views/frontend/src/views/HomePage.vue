<template>
  <div class="app-container">
    <Sidebar />
    <div class="grp-container">
      <div v-if="mode === 'view'" style="margin-bottom: 20px;">
        <el-row>
          <el-col :span="3">
            <h2 style="font-size: 16px; font-weight: bold; color: #132F51;">Personal Information</h2>
          </el-col>
          <el-col :span="1">
            <div style="width: 24px; height: 24px;">
              <div @click="changeMode('edit')" style="width: 16px; height: 16px;"><el-icon>
                  <EditPen />
                </el-icon></div>
            </div>
          </el-col>
        </el-row>
      </div>
      <div v-else-if="mode === 'edit'" style="margin-bottom: 20px;">
        <el-row>
          <el-col :span="4">
            <h2 style="font-size: 16px; font-weight: bold; color: #132F51;">Edit Personal Information</h2>
          </el-col>
          <el-col :span="2">
            <el-button @click="changeMode('view')" type="success">Save</el-button>
          </el-col>
          <el-col :span="2">
            <el-button @click="changeMode('view')" type="danger">Cancel</el-button>
          </el-col>
        </el-row>
      </div>

      <div v-if="mode === 'view'">
        <el-row>
          <div class="grp-center">
            <div style="width: 40px; height: 40px;"><el-icon>
                <User />
              </el-icon></div>
          </div>
        </el-row>
        <el-row>
          <div style="font-size: 14px;"><strong>Personal Details</strong></div>
        </el-row>

        <div v-if="userDetails">
          <div class="content-wrapper">
            <el-row>
              <div class="content">
                <el-row :gutter="10" style="width: 300px;">
                  <el-col :span="15"><label>Firstname</label></el-col>
                  <el-col :span="8"><label>Lastname</label></el-col>
                </el-row>
                <el-row :gutter="10" style="width: 300px;">
                  <el-col :span="15">
                    <p class="font-12">{{ userDetails.firstname }}</p>
                  </el-col>
                  <el-col :span="8">
                    <p class="font-12">{{ userDetails.lastname }}</p>
                  </el-col>
                </el-row>
              </div>
            </el-row>
            <el-row>
              <div class="content">
                <el-row :gutter="10" style="width: 300px;">
                  <el-col :span="15"><label>Email</label></el-col>
                </el-row>
                <el-row :gutter="10" style="width: 300px;">
                  <el-col :span="15">
                    <p class="font-12">{{ userDetails.email }}</p>
                  </el-col>
                </el-row>
              </div>
            </el-row>
          </div>
        </div>
        <div v-else>
          <p>Loading user details...</p>
        </div>
        <div v-if="error">
          <p>Error: {{ error }}</p>
        </div>
      </div>
      <div v-else-if="mode === 'edit'">
        <div class="content">
          <el-row :gutter="10" style="width: 300px;">
            <el-col :span="12"><label>Firstname</label></el-col>
            <el-col :span="12"><label>Lastname</label></el-col>
          </el-row>
          <el-row :gutter="10" style="width: 300px;">
            <el-col :span="12">
              <el-input v-model="userDetails.firstname" class="font-12"/>
            </el-col>
            <el-col :span="12">
              <el-input v-model="userDetails.lastname" class="font-12"/>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
  <router-view></router-view>

</template>

<script>
import axios from 'axios';
import Sidebar from '/src/components/Sidebar.vue';
import { ElRow, ElCol, ElButton, ElInput } from 'element-plus';
import { notification } from "/src/utils/notification";
import Cookies from 'js-cookie';

export default {
  name: 'HomePage',
  components: {
    Sidebar,
    ElRow,
    ElCol,
    ElButton,
    ElInput
  },
  data() {
    return {
      mode: 'view',
      editMode: false,
      userDetails: null,
      error: null
    };
  },
  mounted() {
    this.fetchUserDetails();
  },
  methods: {
    fetchUserDetails() {
      const token = Cookies.get('token');
      console.log(token);
      if (!token) {
        this.$router.push('/');
        return;
      }

      axios.get('http://localhost:8080/api/v1/userDetails', {
        headers: {
          Authorization: `${Cookies.get('token')}`
        }
      })
        .then(response => {
          this.userDetails = response.data.user;

          // save userID in local storage
          localStorage.setItem('currentUserID', response.data.user.ID);
          this.error = null;
        })
        .catch(error => {
          console.error('Error fetching user details:', error.response || error);
          this.error = 'Error fetching user details.';
          this.userDetails = null;
        });
    },

    changeMode(mode) {
      this.mode = mode;
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

.content {
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
  width: 100px;
  height: 100px;
  color: #ffffff;
  border-radius: 100px;
  background-color: #132F51;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>