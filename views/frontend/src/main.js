import { createApp } from 'vue';
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';
import 'element-plus/dist/index.css'
import * as ElIcon from '@element-plus/icons-vue'

import App from './App.vue';
import router from './router';


import '/public/assets/css/base.css'
import '/public/assets/css/base.scss';
import '/public/assets/css/element-variables.scss';
// import 'element-plus/theme-chalk/index.css';


const app = createApp(App);

Object.keys(ElIcon).forEach((key) => {
    app.component(key, ElIcon[key])
})

app.component(ElButton.name, ElButton);
app.component(ElForm.name, ElForm);
app.component(ElFormItem.name, ElFormItem);
app.component(ElInput.name, ElInput);

app.use(router);
app.mount('#app');
