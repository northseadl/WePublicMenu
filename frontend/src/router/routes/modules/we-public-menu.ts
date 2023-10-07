import { AppRouteRecordRaw } from '@/router/routes/types';
import { PLUGIN_LAYOUT } from '@/router/routes/base';

const WePublicMenu: AppRouteRecordRaw = {
  path: '/we-public-menu',
  name: 'WePublicMenu',
  redirect: '/we-public-menu/index',
  component: PLUGIN_LAYOUT,
  meta: {
    locale: 'menu.example',
    icon: 'icon-settings',
    requiresAuth: true,
    order: 3,
  },
  children: [
    {
      path: '/we-public-menu/index',
      name: 'WePublicMenuIndex',
      component: () => import('@/views/we-public-menu/index.vue'),
      meta: {
        locale: 'menu.example',
        requiresAuth: true,
      },
    },
  ],
};

export default WePublicMenu;
