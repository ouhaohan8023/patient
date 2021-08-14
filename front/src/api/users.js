import request from '@/plugins/request';

export function GetUserLists (data) {
    let params = new URLSearchParams(data);
    return request({
        url: 'lists',
        method: 'get',
        params
    });
}

export function UpdateUserStatus (id) {
    return request({
        url: 'lists/' + id,
        method: 'post'
    });
}
