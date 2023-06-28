# gmonitor

一个用来监控自己的香橙派的小工具.因为手边没有可用的屏幕,经常开机后自动连接到了不知道哪个网络而只能依赖串口调试查看IP地址,于是写了这个玩意.

作用:把网络名称和IP地址输出到8080端口,使用浏览器访问内网穿透(例如[cpolar](https://www.cpolar.com/))的URL即可获得.

未来不太可能会加入更多功能,因为太懒了.

```
 ____________________________________
/                                    \
| BUPT-portal                        |
| 192.168.10.18                      |
| cpu_thermal-virtual-0              |
| Adapter: Virtual device            |
| temp1: +56.9°C (crit = +105.0°C) |
|                                    |
| gpu_thermal-virtual-0              |
| Adapter: Virtual device            |
| temp1: +56.6°C                    |
|                                    |
| ddr_thermal-virtual-0              |
| Adapter: Virtual device            |
| temp1: +57.8°C                    |
|                                    |
| ve_thermal-virtual-0               |
| Adapter: Virtual device            |
| temp1: +57.1°C                    |
|                                    |
\                                    /
 ------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```
