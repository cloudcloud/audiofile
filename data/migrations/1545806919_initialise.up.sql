create table directories (
  id text not null primary key,
  directory text not null,
  date_added timestamp,
  date_updated timestamp
);
create unique index dir_idx on directories(directory);

create table artists (
  id text not null primary key,
  name text,
  status text,
  date_added timestamp,
  url_text text
);

create table albums (
  id text not null primary key,
  name text,
  release_year int,
  date_added timestamp,
  track_count int,
  url_text text
);

create table songs (
  id text not null primary key,
  name text,
  running_length int,
  track_number int,
  tag_major int,
  tag_minor int,
  id3_comment text
);

create table album_artists (
  album_id text not null,
  artist_id text not null,
  date_added timestamp,
  date_updated timestamp
);

create table album_songs (
  album_id text not null,
  song_id text not null,
  date_added timestamp,
  date_updated timestamp
);

create table files (
  id text not null primary key,
  filename text not null,
  song_id text not null,
  date_added timestamp,
  date_updated timestamp
);
