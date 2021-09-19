create database if not exists tictactoe;
use tictactoe;

create table if not exists game (
    game_id int primary key auto_increment,
    player_x_alias varchar(15) not null,
    player_y_alias varchar(15) not null,
    winner enum('x', 'y', 'draw') null,
    game_over tinyint(1) default '0',
    created datetime default current_timestamp,
    modified datetime null on update current_timestamp
);

create table if not exists game_move(
    move_id bigint(20) primary key auto_increment,
    game_id int not null,
    player_x_move varchar(2) null default '--',
    player_y_move varchar(2) null default '--',
    created datetime default current_timestamp,
    modified datetime null on update current_timestamp,

    foreign key (game_id) references game(game_id)
);