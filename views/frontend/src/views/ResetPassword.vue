<template>
  <div class="grp-container">
    <div class="grp-form">
      <div class="grp-row">
        <h2
          style="font-size: 16px; font-weight: bold; color: #132F51;"
        >
          Sign Up
        </h2>
      </div>

      <el-form ref="signupForm" :model="form" :rules="rules" label-position="top" style="width: 100%">
        <div style="width: 100%; display: flex; flex-direction: row; justify-content: space-between;">
          <el-form-item label="First Name" prop="firstName" style="width: 45%;">
            <el-input v-model="form.firstName" clearable></el-input>
          </el-form-item>
          <el-form-item label="Last Name" prop="lastName" style="width: 45%;">
            <el-input v-model="form.lastName" clearable></el-input>
          </el-form-item>
        </div>
            
        <el-form-item label="Email" prop="email">
          <el-input v-model="form.email" clearable></el-input>
        </el-form-item>

        <el-form-item label="Password" prop="password">
          <el-input type="password" v-model="form.password"  show-password clearable></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm">Reset</el-button>
          <el-button class="button" @click="cancel">Back</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { ElForm, ElFormItem, ElInput, ElButton } from 'element-plus';
import { notification} from "/src/utils/notification";
import Cookies from 'js-cookie';

export default {
  name: 'SignupForm',
  components: {
    ElForm,
    ElFormItem,
    ElInput,
    ElButton,
  },
  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        password: ''
      },
      rules: {
        firstName: [
          { required: true, message: 'Please enter your first name', trigger: 'blur' },
        ],
        lastName: [
          { required: true, message: 'Please enter your last name', trigger: 'blur' },
        ],
        email: [
          { required: true, message: 'Please enter your email', trigger: 'blur' },
          { type: 'email', message: 'Please enter a valid email address', trigger: ['blur', 'change'] },
        ],
        password: [
          { required: true, message: 'Please enter your password', trigger: 'blur' },
          { min: 6, message: 'Password length should be at least 6 characters', trigger: 'blur' },
        ],
      },
      loginForm: {
        email: '',
        password: ''
      },
      loginRules: {
        email: [
          { required: true, message: 'Please enter your email', trigger: 'blur' },
          { type: 'email', message: 'Please enter a valid email address', trigger: ['blur', 'change'] },
        ],
        password: [
          { required: true, message: 'Please enter your password', trigger: 'blur' },
        ],
      }
    };
  },
  methods: {
    submitForm() {
      
    },
    cancel() {
      this.$router.push('/');
    },
  }
};
</script>
  
<style lang="scss" scoped>

::v-deep {
    .el-input__inner{
      width: 100%;
      height: 30px;
    }

    .el-form-item__error {
      font-size: 10px;
      margin-left: -30px;
    }
}

.button{
  border: 0;
  color: var(--color_text_default);
  background-color: #e9f0f0;
  font-weight: bold;
  font-size: 14px;

  &:hover {
    background-color: #132F51;
    color: #FFFFFF;
  }
}

.grp-form {
  width: 60%;
  height: calc(100% - 40px);
  padding: 20px;
  float: right;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.grp-row {
  margin-bottom: 20px;
}

.grp-link {
  font-size: 12px;
  color: #132F51;
  cursor: pointer;
  text-decoration: underline;
  font-weight: bold;
}

.grp-link:hover {
  color: #3C678A;
}
</style>