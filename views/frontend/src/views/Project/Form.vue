<template>
  <div class="grp-container">
    <div class="grp-form">
      <div class="grp-row">
        <h2
          style="font-size: 16px; font-weight: bold; color: #132F51;"
        >
          Create Project
        </h2>
      </div>

      <el-form ref="signupForm" :model="form" :rules="rules" label-position="top" style="width: 100%">
        <!-- <div style="width: 100%; display: flex; flex-direction: row; justify-content: space-between;">
          <el-form-item label="Prefix" prop="projectPrefix" style="flex: 1;">
            <el-input v-model="form.projectPrefix" clearable></el-input>
          </el-form-item>
          <span style="margin: 0 10px;">/</span>
          <el-form-item label="Project Id" prop="projectSuffix" style="flex: 5;">
            <el-input v-model="form.projectSuffix" clearable></el-input>
          </el-form-item>
        </div> -->

        <div style="width: 100%; display: flex; flex-direction: row; gap: 20px;">
          <el-form-item label="General Id" style="width: 50%;">
            <el-col :span="5">
              <el-form-item>
                <el-input v-model="form.projectPrefix" placeholder="PrE"/>
              </el-form-item>
            </el-col>
            <el-col :span="0.5" style="padding: 0px 10px;">
              <span>/</span>
            </el-col>
            <el-col :span="18">
              <el-form-item>
                <el-input v-model="form.projectSuffix" placeholder="Project ID"/>
              </el-form-item>
            </el-col>
          </el-form-item>

          <el-form-item label="Project Name" style="width: 50%;">
            <el-input v-model="form.projectName" placeholder=" "/>
          </el-form-item>
        </div>
        
        <div style="width: 100%; display: flex; flex-direction: row; gap: 20px;">
          <el-form-item label="Start Date" style="width: 25%;">
            <el-date-picker
              v-model="form.startDate"
              type="date"
              placeholder="Select start date"
              format="DD/MM/YYYY"
              style="width: 100%;"
            />
          </el-form-item>

          <el-form-item label="End Date" style="width: 25%;">
            <el-date-picker
              v-model="form.endDate"
              type="date"
              placeholder="Select end date"
              format="DD/MM/YYYY"
              style="width: 100%;"
            />
          </el-form-item>

          <el-form-item label="Type" style="width: 25%;">
            <el-select v-model="value" placeholder="Select" style="width: 100%">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="Status" style="width: 25%;">
            <el-select v-model="value" placeholder="Select" style="width: 100%">
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                :disabled="item.disabled"
              />
            </el-select>
          </el-form-item>
        </div>

        <el-form-item label="Description" style="width: 100%;">
          <el-input v-model="form.projectDescription" placeholder="Description"/>
        </el-form-item>
            
       
        <el-form-item>
          <el-button type="primary" @click="submitForm">Create</el-button>
          <el-button class="button" @click="cancel">Cancel</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
  
<script>
import { ref } from 'vue';

export default {
  name: 'Form',
  components: {
  },
  data(){
    return {
      form: {
        projectPrefix: 'PRE',
        projectSuffix: '0000001',
        projectType: {},
        projectStatus: {},
        projectDescription: '',
        startDate: '',
        endDate: '',
        projectName: '',
      },
      rules: {
        
      },
    };
  },
  methods: {
    submitForm() {
      
    },
    cancel() {
      this.$router.push('/project');
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
  width: calc(100% - 40px);;
  height: calc(100% - 40px);
  padding: 20px;
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