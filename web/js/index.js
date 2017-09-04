/**
 * Created by Sergey on 09/08/2017.
 */

function $(id) {
    //类型的比较
    return typeof id === 'string' ? document.getElementById(id) : id;
}

//当网页加载完毕时调用
window.onload = function () {
    var lis = $('tab-header').getElementsByTagName('li');
    var content = $('tab-content').getElementsByClassName('dom');

    console.log(lis,content);

    //验证
    if (lis.length !== content.length) return;

    for (var i = 0; i < lis.length; i++){
        var li = lis[i];
        console.log(li);

        li.id = i;

        li.onmousemove = function () {
            for (var j = 0;j < lis.length;j++){
                lis[j].className = '';
                content[j].style.display = 'none';
            }

            this.className = 'selected';

            content[this.id].style.display = 'block';
        }
    }
}