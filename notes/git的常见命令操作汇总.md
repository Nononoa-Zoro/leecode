### git的常见命令操作汇总

|      |                                                              |
| ---- | ------------------------------------------------------------ |
|      | 从远程分支pull代码                                           |
|      | 1.新建一个目录执行命令 git init 初始化为git的工作目录        |
|      | 2.执行命令 git remote add origin git@github.com:XXXX/nothing2.git(远程仓库链接) |
|      | 3.将远程分支pull到本地 git fetch origin dev (dev是远程仓库的分支名) |
|      |                                                              |
|      | 在本地创建分支dev并切换到该分支                              |
|      | git checkout -b dev(本地分支名称) origin/dev(远程分支名称)   |
|      |                                                              |
|      |                                                              |
|      | 从远程分支pull代码到本地                                     |
|      | git pull origin dev                                          |
|      |                                                              |
|      |                                                              |
|      | 查看本地分支                                                 |
|      | git branch                                                   |
|      |                                                              |
|      | 删除本地分支                                                 |
|      | git branch -d dev(本地分支名称)                              |
|      |                                                              |
|      | 查看远程分支                                                 |
|      | git branch -a                                                |
|      |                                                              |
|      | 查看本地分支和关联的远程分支                                 |
|      | git branch -vv                                               |
|      |                                                              |
|      | 创建本地分支                                                 |
|      | git branch test                                              |
|      |                                                              |
|      | 切换本地分支                                                 |
|      | git checkout test                                            |
|      |                                                              |
|      | 创建并切换到本地分支                                         |
|      | git checkout -b test  或者　git switch -c test               |
|      |                                                              |
|      |                                                              |
|      | 将本地修改放到stash中，stash其实是一个栈空间                 |
|      | git stash                                                    |
|      |                                                              |
|      | 查看栈空间有多少记录                                         |
|      | git stash list                                               |
|      |                                                              |
|      | 取出stash中的记录恢复                                        |
|      | git stash pop <stash@{0}>                                    |
|      |                                                              |
|      | 查看本地的所有commit记录                                     |
|      | git log 或者 git log --pretty=oneline                        |
|      |                                                              |
|      | 恢复到指定的commit版本                                       |
|      | git reset --hard <commit-id>                                 |
|      |                                                              |
|      | 恢复到上一个版本                                             |
|      | git reset --hard HEAD^                                       |
|      |                                                              |
|      | 恢复到上上个版本                                             |
|      | git reset --hard HEAD^^                                      |
|      |                                                              |
|      | git reset --hard <commit-id> //回退到之前的某一个commit-id的版本 |
|      | git reset --soft <commit-id> //回退到之前的某一个commit-id的版本，但是代码不会回滚，只有commit-id变为之前的版本 |
|      |                                                              |
|      | 查看所有的命令记录                                           |
|      | git reflog                                                   |
|      |                                                              |
|      | 撤销修改                                                     |
|      | git checkout <filename> //把filename在工作区的修改全部撤销掉 回到最近一次的commit或者是add的记录 |
|      |                                                              |
|      | 用于合并指定分支到当前分支                                   |
|      | git merge <branch-name>                                      |
|      |                                                              |
|      | git pull 和 git fetch                                        |
|      |                                                              |
|      | git pull 的过程可以理解为：                                  |
|      | git fetch origin master //从远程主机的master分支拉取最新内容 |
|      | git merge FETCH_HEAD    //将拉取下来的最新内容合并到当前所在的分支中 |
|      |                                                              |
|      | git fetch是将远程主机的最新内容拉到本地，用户在检查了以后决定是否合并到工作本机分支中。(需要手动执行git merge操作) |
|      |                                                              |
|      | git fetch origin <远程分支名称>:<本地分支名称> //从远程分支拉去并在本地创建一个分支与之关联 |
|      |                                                              |
|      | git推送本地分支到远程分支                                    |
|      | 首先需要pull远程分支代码 git pull origin dev（远程分支名称） |
|      | git add .(提交变更到工作区)                                  |
|      | git commit -m "description" (添加变更的描述信息)             |
|      | git push origin (本地分支名称):(远程分支名称)                |
|      |                                                              |
|      | ##############查看本地tag                                    |
|      | git tag                                                      |
|      |                                                              |
|      | ##############查看远程的所有tag                              |
|      | git ls-remote --tags                                         |
|      |                                                              |
|      | ###################                                          |
|      | 创建并切换到本地分支dev                                      |
|      | git checkout -b dev                                          |
|      |                                                              |
|      | ###################                                          |
|      | 查看当前的本地分支与远程分支的关联关系                       |
|      | git branch -vv                                               |
|      |                                                              |
|      | ###################                                          |
|      | 把本地分支与远程origin的分支进行关联处理(通过 --set-upstream-to 命令) |
|      | git branch --set-upstream-to=origin/dev (将本地分支关联到远程dev分支) |
|      |                                                              |
|      |                                                              |
|      | 解除本地分支和远程分支的关联                                 |
|      | 首先切换到你想解除关联的本地分支                             |
|      | git checkout dev                                             |
|      | git remote remove origin                                     |
|      |                                                              |
|      | 修改本地分支名称                                             |
|      | git branch -m oldname newname                                |
|      |                                                              |
|      |                                                              |



#### 删除远程代码

1. #### 仅仅删除远程分支文件，不删除本地文件

   git rm --cached filename 

   git commit -m "delete remote file filename " 

   git push -u origin master(此处是当前分支的名字)

   （如果待删除的是一个文件夹需要加上 -r）

2. ####  删除本地文件与远程分支文件

   git rm filename 

   git commit -m "delete file filename " 

   git push -u origin master(此处是当前分支的名字)

git push -u origin master 相当于 git push origin master  + git branch --set-upstream-to=origin/master



***\*IDEA\**中使用\**Git\**，文件\**不同颜色\**代表的含义**

- 绿色——已经加入控制暂未提交；
- 红色——未加入版本控制；
- 蓝色——加入，已提交，有改动；
- 白色——加入，已提交，无改动；
- 灰色——版本控制已忽略文件；
- 黄色——被git忽略，不跟踪。



### 修改最近一次commit信息

```gitconsole
git commit --amend
```



查看文件修改

git diff file // 此时还没有commit



查看某次commit的某个文件的提交内容

git show commit-id file



查看两个分支的代码差异

git diff branchA branchB file // 查看两个分支在file文件上的差异



### Git reset

index：暂存区，使用git add 提交文件之后，文件就被提交到了暂存区。

working space：工作区，使用git commit 之后，文件的就被提交到了工作区。

HEAD：指向当前分支版本最新的一个提交。



#### git reset的使用

git reset --option <commit-id or branch>

option可选 soft hard mixed

soft：暂存区和工作区的改动都不会变动的情况下，将HEAD指向指定commit-id。

hard：暂存区和工作区全部修改为指定commit-id的状态。

mixed：也是默认的情况，将HEAD重置为指定commit-id，重置暂存区，但是工作区不会变。
