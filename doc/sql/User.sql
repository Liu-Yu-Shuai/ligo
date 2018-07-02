CREATE TABLE user (
  `id` int not null auto_increment comment '自增id',
  `first_name` VARCHAR(20) not null DEFAULT '',
  `last_name` VARCHAR(20) not null DEFAULT '',
  `updated_at` datetime null default null on update current_timestamp comment '更新时间',
  `created_at` datetime not null default current_timestamp comment '创建时间',
  PRIMARY KEY (`id`)
)