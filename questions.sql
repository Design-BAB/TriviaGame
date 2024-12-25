CREATE TABLE qanda (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question TEXT,
    answer1 TEXT,
    answer2 TEXT,
    answer3 TEXT,
    answer4 TEXT,
    rightanswer INTEGER
);

INSERT INTO qanda (question, answer1, answer2, answer3, answer4, rightanswer)
VALUES ("Who was the first person killed?", "Adam", "Eve", "Abel", "Seth", 2);

INSERT INTO qanda (question, answer1, answer2, answer3, answer4, rightanswer)
VALUES ("The Anime Rounin Kenshin and Golden Kamuy take place in an era of Japanese history also known as the ___ era", 
        "Meiji", "Showa", "Edo", "Boshin War", 0);

SELECT * FROM qanda;
