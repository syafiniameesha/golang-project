<template>
  <div class="grp-container">
    <div class="grp-form">
      <div class="grp-row">
        <h2
          style="font-size: 16px; font-weight: bold; color: #132F51;"
        >
          Reset Password
        </h2>
      </div>

      <el-form ref="resetForm" :model="form" :rules="rules" label-position="top" style="width: 100%">
        <el-form-item label="Password" prop="password">
          <el-input type="password" v-model="form.password"  show-password clearable></el-input>
        </el-form-item>

        <el-form-item label="Confirm Password" prop="password">
          <el-input type="password" v-model="form.confirmPassword"  show-password clearable></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm">Reset</el-button>
          <el-button class="button" @click="cancel">Cancel</el-button>
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
        password: '',
        confirmPassword: '',
      },
      rules: {
        password: [
          { required: true, message: 'Please enter your password', trigger: 'blur' },
          { min: 6, message: 'Password length should be at least 6 characters', trigger: 'blur' },
        ],
        confirmPassword: [
          { required: true, message: 'Please enter your password', trigger: 'blur' },
          { min: 6, message: 'Password length should be at least 6 characters', trigger: 'blur' },
        ],
      },
    };
  },
  methods: {
    submitForm() {
      this.$refs.resetForm.validate((valid) => {
        if (valid) {
          const token = this.$route.query.token;
          axios.post(`http://localhost:8080/api/v1/resetPassword/${token}`, {
            password: this.form.password,
            confirmPassword: this.form.confirmPassword,
          })
          .then(response => {
            notification('Password Reset', 'Successfully Reset Your Password', 'bottom-left');
            this.$router.push('/');
          })
          .catch(error => {
            notification('Error Reset Password', '', 'bottom-left');
          });
        } else {
          return false;
        }
      });
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