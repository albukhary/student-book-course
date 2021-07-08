INSERT INTO student (first_name, last_name, email) VALUES ('Lazizbek', 'Kahramonov', 'lazizbek@novalab.uz');
INSERT INTO student (first_name, last_name, email) VALUES ('Khurshidbek', 'Kobilov', 'khurshidbek@mitc.uz');
INSERT INTO student (first_name, last_name, email) VALUES ('Jaloliddin', 'Abdullaev', 'jaloliddin@andriod.uz');
INSERT INTO student (first_name, last_name, email) VALUES ('Abdulaziz', 'Musaev', 'abdulaziz@mitc.uz');
INSERT INTO student (first_name, last_name, email) VALUES ('Uktamjon', 'Abduraupov', 'uktamjon@akfa.uz');

INSERT INTO book (title, author) VALUES ('Tafsiri Hilol', 'Shayx Muhammad Sodiq Muhammad Yusuf');
INSERT INTO book (title, author) VALUES ('Baxtiyor Oila', 'Shayx Muhammad Sodiq Muhammad Yusuf');
INSERT INTO book (title, author) VALUES ( 'Ruhiy tarbiya', 'Shayx Muhammad Sodiq Muhammad Yusuf');
INSERT INTO book (title, author) VALUES ('Olim, Odam va Olam', 'Mubashshir Ahmad');
INSERT INTO book (title, author) VALUES ('Ikki eshik orasi', 'O`tkir Hoshimov');
INSERT INTO book (title, author) VALUES ('Daftar Hoshiyasidagi bitiklar', 'O`tkir Hoshimov');
INSERT INTO book (title, author) VALUES ('O`tkan kunlar', 'Abdulla Qodiriy');\
INSERT INTO book (title, author) VALUES ('Xamsa', 'Alisher Navoiy');

INSERT INTO course (instructor, title) VALUES ('Imam Moturudiy', 'Aqidah');
INSERT INTO course (instructor, title) VALUES ('Imam Abu Hanifa', 'Usul al-fiqh');
INSERT INTO course (instructor, title) VALUES ('Sarvar Abdullaev', 'OOP');
INSERT INTO course (instructor, title) VALUES ( 'Ashish Seth', 'Computer Networks');
INSERT INTO course (instructor, title) VALUES ('Intro to Economics', 'Jong Song Lee');
INSERT INTO course (instructor, title) VALUES ('Engineering Mathematics', 'Anna Tomskova');
INSERT INTO course (instructor, title) VALUES ('Probability and Statistics', 'Bahodir Ahmedov');
INSERT INTO course (instructor, title) VALUES ('Engineering Communications', 'Sejin Kang');

INSERT INTO student_course (student_id, course_id) VALUES (1, 1);
INSERT INTO student_course (student_id, course_id) VALUES (1, 2);
INSERT INTO student_course (student_id, course_id) VALUES (1, 3);

INSERT INTO student_course (student_id, course_id) VALUES (2, 4);
INSERT INTO student_course (student_id, course_id) VALUES (2, 5);
INSERT INTO student_course (student_id, course_id) VALUES (2, 6);

INSERT INTO student_course (student_id, course_id) VALUES (3, 6);
INSERT INTO student_course (student_id, course_id) VALUES (3, 7);
INSERT INTO student_course (student_id, course_id) VALUES (3, 8);

INSERT INTO student_course (student_id, course_id) VALUES (4, 1);
INSERT INTO student_course (student_id, course_id) VALUES (4, 8);

INSERT INTO student_course (student_id, course_id) VALUES (5, 2);
INSERT INTO student_course (student_id, course_id) VALUES (5, 7);

INSERT INTO student_book (student_id, book_id) VALUES (1, 1);
INSERT INTO student_book (student_id, book_id) VALUES (1, 2);
INSERT INTO student_book (student_id, book_id) VALUES (1, 3);

INSERT INTO student_book (student_id, book_id) VALUES (2, 4);
INSERT INTO student_book (student_id, book_id) VALUES (2, 5);
INSERT INTO student_book (student_id, book_id) VALUES (2, 6);

INSERT INTO student_book (student_id, book_id) VALUES (3, 6);
INSERT INTO student_book (student_id, book_id) VALUES (3, 7);
INSERT INTO student_book (student_id, book_id) VALUES (3, 8);

INSERT INTO student_book (student_id, book_id) VALUES (4, 1);
INSERT INTO student_book (student_id, book_id) VALUES (4, 8);

INSERT INTO student_book (student_id, book_id) VALUES (5, 2);
INSERT INTO student_book (student_id, book_id) VALUES (5, 7);