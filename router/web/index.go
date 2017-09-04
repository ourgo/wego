package web

import (
	"net/http"
	"html/template"
	"fmt"
)

func Index(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte(`
	<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.0//EN" "http://www.wapforum.org/DTD/xhtml-mobile10.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>无标题文档</title>
</head>

<body>
<div align="center">
  <table width="382" border="1" align="center" bordercolor="#CC6633">
    <tr>
      <th width="372" height="23" align="left" valign="middle" bgcolor="#009999" scope="col"><form id="form1" method="post" action="">
        <p>
          <label>
            <select name="select" dir="rtl">
              <option>设置</option>
              <option>添加好友</option>
              <option>切换账号</option>
              <option>退出账号</option>
              <option>关闭Wego</option>
            </select>
            </label>
        </p>
        </form>    </th>
    </tr>
  </table>
</div>
<div align="center">
  <table width="382" border="1" align="center">
    <tr>
      <th width="296" height="341" align="left" valign="top" scope="col"><textarea name="textarea" cols="32" rows="23"></textarea></th>
      <th width="70" align="left" valign="top" scope="col"><form id="form2" method="post" action="">
        <p>
        <label>
          <div align="left">
          好友列表

          <select name="select2" size="23">
            <option>好友1</option>
            <option>好友2</option>
            <option>好友3</option>
            <option>好友4</option>
            <option>好友5</option>
          </select>
        </label>
      </form>      </th>
    </tr>
    <tr>
      <td height="72"><textarea name="textarea2" cols="32" rows="3"></textarea></td>
      <td align="center" valign="middle">        <p>
          <label></label>
      </p>
        <form id="form3" method="post" action="">
          <p align="center">
            <label>
              <input type="submit" name="Submit" value="发送" />
            </label>
          </p>
        </form>      </td>
    </tr>
</table>
  <div align="center"><div align="justify"><div align="center"><p align="center"></p>
</div></div></div></body>
</html>
        `))
}
