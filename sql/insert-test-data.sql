insert into oauth_access_tokens(
    id
    , user_id
    , client_id
    , name
    , scopes
    , revoked
    , created_at
    , updated_at
    , expires_at
)
select
    concat('tokentoken', @id)
    , @id := @id + 1
    , 1
    , 'name'
    , 'scopes'
    , 0
    , now()
    , now()
    , now()
from
    (SELECT @id := 1000) AS t,
    users as t1,
    users as t2,
    users as t3
limit 1000
;
