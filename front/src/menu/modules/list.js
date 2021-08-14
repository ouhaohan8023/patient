const pre = '/list/';

export default {
    path: '/list',
    title: 'Patients',
    header: 'home',
    custom: 'i-icon-demo i-icon-demo-list',
    children: [
        {
            path: `${pre}table-list`,
            title: 'Lists'
        }
    ]
}
