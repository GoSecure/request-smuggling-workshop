DROP TABLE IF EXISTS movies CASCADE;
CREATE TABLE movies
(
     id INTEGER,
     title      VARCHAR(100), 
     cover        VARCHAR(100), 
     synopsis        VARCHAR(500),
     genre   VARCHAR(50),
     year   INTEGER,
     public   BOOLEAN,
     PRIMARY KEY  (id)
);

INSERT INTO movies VALUES
('1', 'Spiderman', 'tt0145487.jpg', 'Peter Parker\'s life changes when he is bitten by a genetically altered spider and gains superpowers. He uses his powers to help people and finds himself facing the Green Goblin, an evil maniac.', 'action', '2002',true),
('2', 'Matrix', 'tt0234215.jpg', 'Thomas Anderson, a computer programmer, is led to fight an underground war against powerful computers who have constructed his entire reality with a system called the Matrix.', 'science fiction', '1999',true),
('3', 'Dark Knight', 'tt0468569.jpg', 'After Gordon, Dent and Batman begin an assault on Gotham\'s organised crime, the mobs hire the Joker, a psychopathic criminal mastermind who offers to kill Batman and bring the city to its knees.', 'action', '2008',true),
('4', 'Tron', 'tt1104001.jpg', 'Sam misses his father, a virtual world designer, and enters a virtual space that has become much more dangerous than his father intended. Now, both father and son embark upon a life-and-death journey.', 'science fiction', '2010',true),
('5', 'Inception', 'tt1375666.jpg', 'Cobb steals information from his targets by entering their dreams. Saito offers to wipe clean Cobb\'s criminal history as payment for performing an inception on his sick competitor\'s son.', 'suspence', '2010',true),
('6', 'Dune', 'tt202111.jpg', 'Paul Atreides, a brilliant and gifted young man born into a great destiny beyond his understanding, must travel to the most dangerous planet in the universe to ensure the future of his family and his people. As malevolent forces explode into conflict over the planet\'s exclusive supply of the most precious resource in existence, only those who can conquer their own fear will survive. FLAG-629715B21387E95994E5FB725CBD4BD8', 'science fiction', '2021',false)

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users
(
     id INTEGER,
     username VARCHAR(100), 
     password VARCHAR(100),
     PRIMARY KEY  (id)
);
INSERT INTO users VALUES
('1', 'admin', 'FLAG-5CDC2509FE6C98B80DCCD578DDDE8CE1'),
('2', 'phil', 'FLAG-11A144C4059E9B7F090FC1D195B8E274')
