极客时间错误处理作业提交

问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么？

答：不应该Wrap抛给上层，上层应该只需要处理Dao层的错误，而不应该接触到底层DB的错误。