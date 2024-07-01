import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '',
    redirect: '/signup',
    children: [
      {
        path: '/signup',
        name: 'Signup',
        component:() => import('/src/views/SignupForm.vue'),
      }
    ]
  },
  {
    path: '/homepage',
    name: 'Home',
    component:() => import('/src/views/HomePage.vue'),
  },
  {
    path: '/resetPassword',
    name: 'Reset Password',
    component:() => import('/src/views/ResetPassword.vue'),
  },
  {
    path: '/project',
    name: 'Project',
    component:() => import('/src/views/Project/Project.vue'),
    children: [
      {
        path: 'form',
        component: () => import('/src/views/Project/Form.vue'),
        name: 'ProjectForm',
        props: true,
      },
    ]
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
