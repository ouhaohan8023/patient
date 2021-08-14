import request from '@/plugins/request';

export function AccountLogin (data) {
    return request({
        url: '/login',
        method: 'post',
        data
    });
}

export function AccountRegister (data) {
    return request({
        url: '/submit',
        method: 'post',
        data
    });
}
