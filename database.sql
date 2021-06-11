select u.ID, u.UserName, p.ParentUserName from "USER" u, "PARENT" p
where u.Parent = p.ID

select u.ID, u.UserName, p.ParentUserName from "USER" u
inner join "PARENT" p
on u.Parent = p.ID