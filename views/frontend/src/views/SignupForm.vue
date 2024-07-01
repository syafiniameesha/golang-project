<template>
  <div class="grp-container">
    <!-- grp logo -->
     <div class="grp-logo">
      <img src="/public/logo.svg" alt=""/>
     </div>
    <!-- signup -->
    <div class="grp-form" v-if="view === 'signup'">
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
          <el-button type="primary" @click="submitForm">Submit</el-button>
          <el-button class="button" @click="updateView('login')">Login</el-button>
        </el-form-item>
      </el-form>

      <el-link type="primary"
        class="grp-link"
        style=""
        @click="updateView('forgotPassword')"
      >
        Forgot Password
      </el-link>
    </div>

    <!-- login -->
    <div v-else-if="view === 'login'" class="grp-form">
      <div class="grp-row">
        <h2 style="font-size: 16px; font-weight: bold; color: #132F51;">Login</h2>
      </div>
      <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-position="top" style="width: 100%">
        <div class="grp-row">
          <el-form-item label="Email" prop="email">
            <el-input v-model="loginForm.email" clearable></el-input>
          </el-form-item>
        </div>

        <div class="grp-row">
          <el-form-item label="Password" prop="password">
            <el-input type="password" v-model="loginForm.password" show-password clearable></el-input>
          </el-form-item>
        </div>

        <div class="grp-row">
          <el-form-item>
            <el-button type="primary" @click="login">Login</el-button>
            <el-button class="button" @click="updateView('signup')">Signup</el-button>
          </el-form-item>
        </div>

        <el-link type="primary"
          class="grp-link"
          style=""
          @click="updateView('forgotPassword')"
        >
          Forgot Password
        </el-link>
      </el-form>
    </div>

    <!-- forgot Password -->
    <div v-else-if="view === 'forgotPassword'" class="grp-form">
      <div class="grp-row">
        <h2 style="font-size: 16px; font-weight: bold; color: #132F51;">Forgot Password</h2>
      </div>
      <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-position="top" style="width: 100%">
        <div class="grp-row">
          <el-form-item label="Email" prop="email">
            <el-input v-model="loginForm.email" clearable></el-input>
          </el-form-item>
        </div>

        <div class="grp-row">
          <el-form-item>
            <el-button type="primary" @click="forgotPassword">Send</el-button>
            <el-button class="button" @click="updateView('signup')">Cancel</el-button>
          </el-form-item>
        </div>
      </el-form>
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { notification} from "/src/utils/notification";
import Cookies from 'js-cookie';

export default {
  name: 'SignupForm',
  components: {
  },
  data() {
    return {
      tokenExpirationTimer: null,
      tokenExpirationTime: 60 * 60 * 1000, //sey expiration time in seconds
      view: 'signup',
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
      const formData = {
        firstName: this.form.firstName,
        lastName: this.form.lastName,
        email: this.form.email,
        password: this.form.password
      };

      axios.post('http://localhost:8080/api/v1/signup', formData)

      .then(response => {
        const token = response.data.token;
        localStorage.setItem('token', token);
        notification('Account Created','Account Created Successfully', 'bottom-left');
        this.view = 'login';
      })

      .catch(error => {
      });
    },
    updateView(view) {
      this.view = view;
    },
    login() {
      const formData = {
        email: this.loginForm.email,
        password: this.loginForm.password
      };

      axios.post('http://localhost:8080/api/v1/login', formData)
      .then(response => {
        const token = response.data.token;
        const fullname = response.data.user.firstname + " " + response.data.user.lastname;
        Cookies.set('token', token, { expires: new Date(Date.now() + this.tokenExpirationTime) });
        this.startTokenExpirationTimer();
        this.$router.push('/homepage'); // Redirect to homepage after login
        notification('Welcome ' + fullname, 'Successfully Logged In', 'bottom-left');
      })
      .catch(error => {
      });
    },
    forgotPassword() {
      const formData = {
        email: this.loginForm.email,
      };

      axios.post('http://localhost:8080/api/v1/forgotPassword', formData)
      .then(response => {
        this.$router.push('/'); 
        notification('Forgot Password', 'Forgot Password Link Send to Your Email', 'bottom-left');
      })
      .catch(error => {
      });
    },
    startTokenExpirationTimer() {
      this.clearTokenExpirationTimer();
      this.tokenExpirationTimer = setTimeout(() => {
        this.clearToken();
        this.$router.push('/');
        notification('Session Expired', 'Please login again', 'bottom-left');
      }, this.tokenExpirationTime);
    },
    clearTokenExpirationTimer() {
      if (this.tokenExpirationTimer) {
        clearTimeout(this.tokenExpirationTimer);
        this.tokenExpirationTimer = null;
      }
    },
    clearToken() {
      Cookies.remove('token');
    }
  },
  created() {
    const token = Cookies.get('token');
    if (token) {
      this.startTokenExpirationTimer();
    } else {
      this.$router.push('/');
    }
  },
  beforeDestroy() {
    this.clearTokenExpirationTimer();
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

.grp-logo {
  float: left;
  width: 35%;
  height: 100%;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;

  img {
    width: 70%;
    height: 70%;
    object-fit: contain;
  }

  @media screen and (max-width: 830px) {
    width: 100%;
    height: 30%;
    max-width: 100%; 
    padding-left: 40px;
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