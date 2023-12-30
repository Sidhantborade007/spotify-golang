CREATE DATABASE SpotifyDB;
USE SpotifyDB;


DROP TABLE IF EXISTS `tracks`;
CREATE TABLE `tracks` (
  `track_name` varchar(200) NOT NULL,
  `album_release_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `album_name` varchar(200) NOT NULL,
  `track_number` int(11) NOT NULL,
  `popularity` int(11) NOT NULL,
  `id` varchar(100) NOT NULL,
  `isrc` varchar(100) NOT NULL,
  PRIMARY KEY (`isrc`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tracks` VALUES
	('Goodbye Everyone','1978-12-15','Alive', 13 ,10, '5tShpGUKu8hKAIBSJUEbB0', 'JPJ931400245'),
	('story 35','2023-11-10','The Count Down', 13 ,17, '6UJLH0ligpcBA1KXMoh0vt','TCAHP2347951'),
	('2 Die 4','2023-11-10','The Count Down', 15 ,11, '5xEb2HT8XoIonmgBvkQShF', 'TCAHP2347902'),
	('GOODBYE - Slowed','1987-12-26','Alive', 13 ,10, '5g8asdwVwYYbSmPaIDEvPj', 'QZHNA2339649');
    
    select * from tracks;
    
    
DROP TABLE IF EXISTS `artist`;
    CREATE TABLE `artist` (
      `id` varchar(100) NOT NULL,
      `href` varchar(200) NOT NULL,
      `name` varchar(100) NOT NULL,
       `uri` varchar(200) NOT NULL,
       `isrc` varchar(100) NOT NULL,
      KEY `artist_FK` (`isrc`),
      CONSTRAINT `artist_FK` FOREIGN KEY (`isrc`) REFERENCES `tracks` (`isrc`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `artist` VALUES
    (  '4qfQTMwNNuBKJ4rF9Lo4bD', 'https://api.spotify.com/v1/artists/4qfQTMwNNuBKJ4rF9Lo4bD', 'Siinamota', 'spotify:artist:4qfQTMwNNuBKJ4rF9Lo4bD', 'JPJ931400245'),
    (  '3vpv7sEkwDmtjDVP3Q0bXH', 'https://api.spotify.com/v1/artists/3vpv7sEkwDmtjDVP3Q0bXH', 'Go Yayo', 'spotify:artist:3vpv7sEkwDmtjDVP3Q0bXH', 'TCAHP2347951'),
    (  '4qfQTMwNNuBKJ4rF9Lo4bD', 'https://api.spotify.com/v1/artists/4qfQTMwNNuBKJ4rF9Lo4bD', 'Go Yayo', 'spotify:artist:3vpv7sEkwDmtjDVP3Q0bXH', 'TCAHP2347951'),
    (  '5X3GK5ziGzILgGeE4b4I97', 'https://api.spotify.com/v1/artists/5X3GK5ziGzILgGeE4b4I97', 'Bedo', 'spotify:artist:5X3GK5ziGzILgGeE4b4I97', 'QZHNA2339649');
      
select * from artist
    
    
    
    