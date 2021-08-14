import BasicLayout from '@/layouts/basic-layout';

const meta = {
    auth: true
};

const pre = 'list-';

export default {
    path: '/list',
    name: 'list',
    redirect: {
        name: `${pre}table-list`
    },
    meta,
    component: BasicLayout,
    children: [
        {
            path: 'table-list',
            name: `${pre}table-list`,
            meta: {
                ...meta,
                title: 'Patient Lists'
            },
            component: () => import('@/pages/list/table-list')
        }
    ]
};
