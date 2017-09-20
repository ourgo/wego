/**
 * Created by Sergey on 09/08/2017.
 */

function $(id) {
    return typeof id === 'string' ? document.getElementById(id) : id;
}
