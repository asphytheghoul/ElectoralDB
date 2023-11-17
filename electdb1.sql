create database electdb;
use electdb;

create table constituency ( constituency_id int primary key, constituency_name varchar(30) NOT NULL, voter_count int default 0, unique key (constituency_name));

create table poll_booth (poll_booth_id int primary key, poll_booth_address varchar(30) NOT NULL, voter_count int default 0, evm_vvpat_no int unique NOT NULL ,constituency_id int NOT NULL, foreign key (constituency_id) references constituency(constituency_id));

alter table constituency add state varchar(30); 

create table voter ( aadhar_id int unique, first_name varchar(30) NOT NULL, last_name varchar(30) NOT NULL, middle_name varchar(30) default NULL, gender char(6) NOT NULL, dob date NOT NULL, age int, state varchar(30) NOT NULL, phone_no int NOT NULL, constituency_name varchar(30), foreign key (constituency_name) references constituency(constituency_name), poll_booth_id int NOT NULL, foreign key (poll_booth_id) references poll_booth(poll_booth_id), voter_id varchar(20) primary key);

delimiter //
create trigger set_voter_id before insert on voter for each row begin set NEW.voter_id = CONCAT(NEW.aadhar_id,NEW.poll_booth_id); end// 
delimiter ;

delimiter //
create trigger voter_age before insert on voter for each row begin set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); END;//

set global event_scheduler = ON//

create event update_age on schedule every 1 month do begin update voter set age = TIMESTAMPDIFF(YEAR, dob, CURDATE()); END;//

create trigger voter_age_on_update before update on voter for each row begin if NEW.dob != OLD.dob then set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); end if; end//

create trigger prevent_delete_constituency before delete on constituency for each row begin declare voters_count int; select count(*) into voters_count from voter where voter.constituency_name = OLD.constituency_name; if voters_count>0 then signal SQLSTATE '45000' set message_text = 'Cannot delete constituency when voters present in it'; end if; end//

create trigger prevent_delete_pollbooth before delete on poll_booth for each row begin declare voters_count int; select count(*) into voters_count from voter where voter.poll_booth_id = OLD.poll_booth_id; if voters_count>0 then signal SQLSTATE '45000' set message_text = 'Cannot delete poll booth when voters present in it'; end if; end//

delimiter ;

create table election (election_id int primary key, election_type char(20) NOT NULL, seats int default 0, dateofelection date, winner char(20));

create table election_cons (elec_id int, cons_id int, foreign key (elec_id) references election(election_id), foreign key (cons_id) references constituency(constituency_id) , primary key (elec_id,cons_id));

alter table election_cons add winner_c int;

delimiter //
create trigger update_seats after insert on election_cons for each row begin update election set election.seats = election.seats + 1 where election.election_id = NEW.elec_id; end;//
delimiter ;

create table official(aadhar_id int unique, first_name varchar(30) NOT NULL, last_name varchar(30) NOT NULL, middle_name varchar(30) default NULL, gender char(6) NOT NULL, dob date NOT NULL, age int, phone_no int NOT NULL, constituency_assigned varchar(30), foreign key (constituency_assigned) references constituency(constituency_name), poll_booth_assigned int, foreign key (poll_booth_assigned) references poll_booth(poll_booth_id), official_id varchar(20) primary key);

delimiter //
create trigger set_official_id before insert on official for each row begin set NEW.official_id = CONCAT(NEW.aadhar_id,'off'); end// 
delimiter ;

delimiter //
create trigger off_age before insert on official for each row begin set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); END;//

set global event_scheduler = ON//

create event update_off_age on schedule every 1 month do begin update official set age = TIMESTAMPDIFF(YEAR, dob, CURDATE()); END;//

create trigger official_age_on_update before update on official for each row begin if NEW.dob != OLD.dob then set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); end if; end//
delimiter ;

alter table official add official_rank char(25) NOT NULL;

alter table official add higher_rank_id int; 
alter table official modify column higher_rank_id varchar(20);
alter table official add foreign key (higher_rank_id) references official(official_id); 

create table party(party_name char(25) primary key, party_symbol char(10) unique, president char(30), party_funds int, headquarters char(20), seats_won int, party_leader varchar(20));

create table candidate(aadhar_id int unique, first_name varchar(30) NOT NULL, last_name varchar(30) NOT NULL, middle_name varchar(30) default NULL, gender char(6) NOT NULL, dob date NOT NULL, age int, phone_no int NOT NULL,cons_fight varchar(30), candidate_id varchar(20) primary key, foreign key (cons_fight) references constituency(constituency_name));


delimiter //
create trigger set_canidate_id before insert on candidate for each row begin set NEW.candidate_id = CONCAT(NEW.aadhar_id,'can'); end// 
delimiter ;

delimiter //
create trigger cand_age before insert on candidate for each row begin set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); END;//

set global event_scheduler = ON//

create event update_cand_age on schedule every 1 month do begin update candidate set age = TIMESTAMPDIFF(YEAR, dob, CURDATE()); END;//

create trigger cand_age_on_update before update on candidate for each row begin if NEW.dob != OLD.dob then set NEW.age = TIMESTAMPDIFF(YEAR, NEW.dob, CURDATE()); end if; end//
delimiter ;

alter table election_cons modify winner_c varchar(20);

alter table election_cons add foreign key (winner_c) references candidate(candidate_id);

alter table voter modify constituency_name varchar(30) NOT NULL;

alter table candidate add column party_rep char(25) NOT NULL, add foreign key (party_rep) references party(party_name);

alter table party add column party_member_count int default 0;

delimiter //
create trigger add_count after insert on voter for each row begin update poll_booth set poll_booth.voter_count = poll_booth.voter_count +1 where poll_booth.poll_booth_id = NEW.poll_booth_id; update constituency set constituency.voter_count = constituency.voter_count+1 where constituency.constituency_name =  NEW.constituency_name; end//
delimiter ;

delimiter //
create trigger sub_count after delete on voter for each row begin update poll_booth set poll_booth.voter_count = poll_booth.voter_count -1 where poll_booth.poll_booth_id = OLD.poll_booth_id; update constituency set constituency.voter_count = constituency.voter_count-1 where constituency.constituency_name =  OLD.constituency_name; end//
delimiter ;

drop table election_cons;

alter table constituency add column election_id int, add foreign key(election_id) references election(election_id);
alter table constituency add column current_mla varchar(20) unique, add foreign key (current_mla) references candidate(candidate_id);

alter table voter modify phone_no varchar(10);
alter table voter modify aadhar_id varchar(5);
alter table candidate modify phone_no varchar(10);
alter table candidate modify aadhar_id varchar(5);
alter table official modify phone_no varchar(10);
alter table official modify aadhar_id varchar(5);

insert into constituency(constituency_id,constituency_name,state) values(1,'Bangalore Urban','Karnataka') ;
insert into constituency(constituency_id,constituency_name,state) values(2,'Bangalore Rural','Karnataka');
insert into constituency(constituency_id,constituency_name,state) values(3,'Belagavi','Karnataka'); 
insert into constituency(constituency_id,constituency_name,state) values(4,'Mandya','Karnataka'); 
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(1,'Indiranagar',1,1);
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(2,'K R Pura',12,1); 
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(3,'Jaya Nagar',25,1); 

insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(4,'Bidadi',30,2); 
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(5,'Anekal',31,2); 

insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(6,'Gokak',66,3);
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(7,'Ramdurg',60,3);  

insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(8,'Malavalli',5,4); 
insert into poll_booth(poll_booth_id,poll_booth_address,evm_vvpat_no,constituency_id) values(9,'Maddur',8,4); 

insert into voter(aadhar_id,first_name,last_name,middle_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('9269','Revanth','Ambati','Sreeram','Male','2003-11-22','Karnataka','9663510096','Bangalore Urban',1);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('6047','Akash','Kamalesh','Male','2003-10-29','Karnataka','9876543211','Bangalore Urban',1);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('4772','Aditi','Prabhu','Female','2003-06-16','Karnataka','9876258363','Bangalore Rural',4);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('2589','Abhignya','Kotha','Female','2004-03-05','Karnataka','9876578965','Bangalore Rural',5);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('4528','Adithya','B','Male','2003-01-14','Karnataka','9876456789','Mandya',9);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('9654','Adithya','Korishetty','Male','2003-12-12','Karnataka','9872583691','Belagavi',7);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('6047','Akash','Kamalesh','Male','2003-10-29','Karnataka','9876543211','Bangalore Urban',1);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('8527','Anirudh','Joshi','Male','2003-09-09','Karnataka','9876554321','Belagavi',6);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('5239','Amish','Gupta','Male','2003-09-29','Karnataka','9876556974','Mandya',8);
insert into voter(aadhar_id,first_name,last_name,gender,dob,state,phone_no,constituency_name,poll_booth_id) values ('7539','Aayush','Nagar','Male','2002-03-21','Karnataka','9632581478','Bangalore Urban',3);

insert into party(party_name,party_symbol,president,party_funds,headquarters) values('Bharatiya Janata Party','Lotus','J P Nadda',5000,'New Delhi');
insert into party(party_name,party_symbol,president,party_funds,headquarters) values('Aam Aadmi Party','Broom','Arvind Kejirwal',2500,'New Delhi');
insert into party(party_name,party_symbol,president,party_funds,headquarters) values('Indian National Congress','Palm','Mallikarjun Kharge',3300,'New Delhi');
insert into party(party_name,party_symbol,president,party_funds,headquarters) values('Janata Dal','Farmer','H D Devegowda',800,'Bangalore');

delimiter //
create trigger add_party after insert on candidate for each row begin update party set party.party_member_count = party.party_member_count +1 where party.party_name = NEW.party_rep; end//
delimiter ;

delimiter //
create trigger sub_party after delete on candidate for each row begin update party set party.party_member_count = party.party_member_count -1 where party.party_name = OLD.party_rep; end//
delimiter ;

INSERT INTO candidate (aadhar_id, first_name, last_name, gender, dob, phone_no, cons_fight, party_rep) 
VALUES 
    ('12345', 'Aarav', 'Gupta', 'Male', '1990-01-01', '1234567890', 'Bangalore Urban', 'Bharatiya Janata Party'),
    ('67890', 'Aditi', 'Sharma', 'Female', '1988-05-10', '2345678901', 'Bangalore Urban', 'Indian National Congress'),
    ('23456', 'Arjun', 'Patel', 'Male', '1987-09-20', '3456789012', 'Bangalore Urban', 'Janata Dal'),
    ('78901', 'Dia', 'Singh', 'Female', '1986-11-15', '4567890123', 'Mandya', 'Aam Aadmi Party'),
    ('34567', 'Kabir', 'Yadav', 'Male', '1985-03-25', '5678901234', 'Mandya', 'Bharatiya Janata Party'),
    ('89012', 'Isha', 'Kaur', 'Female', '1984-07-08', '6789012345', 'Belagavi', 'Indian National Congress'),
    ('45678', 'Mira', 'Reddy', 'Male', '1983-12-30', '7890123456', 'Bangalore Rural', 'Janata Dal'),
    ('90123', 'Neha', 'Mishra', 'Female', '1982-04-14', '8901234567', 'Belagavi', 'Aam Aadmi Party'),
    ('56789', 'Rohan', 'Joshi', 'Male', '1981-08-05', '9012345678', 'Bangalore Rural', 'Bharatiya Janata Party'),
    ('01234', 'Sara', 'Kulkarni', 'Female', '1980-10-28', '0123456789', 'Bangalore Rural', 'Indian National Congress');
    
insert into election(election_id,election_type,dateofelection) values(1,'Lok Sabha','2023-11-17');
insert into election(election_id,election_type,dateofelection) values(2,'Assembly','2024-01-22');

alter table official modify official_rank varchar(40);

INSERT INTO official (aadhar_id, first_name, last_name, middle_name, gender, dob, phone_no, constituency_assigned, poll_booth_assigned, official_rank,higher_rank_id) VALUES
('12345', 'Aarav', 'Sharma', 'Kumar', 'Male', '1990-05-12', '9876543210', NULL, NULL, 'Chief Election Commissioner',NULL),
('23456', 'Diya', 'Patel', 'Nisha', 'Female', '1985-08-24', '8765432109', NULL, NULL, 'Election Commissioner','12345off'),
('34567', 'Ananya', 'Gupta', 'Singh', 'Female', '1982-11-03', '7654321098', 'Bangalore Urban', NULL, 'Electoral Officer','23456off'),
('45678', 'Krish', 'Kumar', 'Reddy', 'Male', '1976-04-17', '6543210987', 'Belagavi', NULL, 'Electoral Officer','23456off'),
('56789', 'Rhea', 'Shah', 'Iyer', 'Female', '1988-09-30', '5432109876', 'Bangalore Rural', NULL, 'Electoral Officer','23456off'),
('67890', 'Aryan', 'Chatterjee', 'Sinha', 'Male', '1995-02-08', '4321098765', 'Mandya', NULL, 'Electoral Officerr','23456off'),
('78901', 'Meera', 'Deshmukh', 'Kapoor', 'Female', '1979-07-21', '3210987654', 'Bangalore Urban', 1, 'Poll Booth Worker','34567off'),
('89012', 'Arjun', 'Tiwari', 'Joshi', 'Male', '1987-12-15', '2109876543', 'Belagavi', 6, 'Poll Booth Worker','45678off'),
('90123', 'Niharika', 'Singhania', 'Gupta', 'Female', '1984-06-29', '1098765432', 'Bangalore Rural', 4, 'Poll Booth Worker','56789off'),
('01234', 'Vikram', 'Reddy', 'Menon', 'Male', '1980-03-05', '9876543210', 'Mandya', 8, 'Poll Booth Worker','67890off'),
('12341', 'Kavya', 'Mishra', 'Ahuja', 'Female', '1992-10-18', '8765432109', 'Bangalore Urban', 2, 'Poll Booth Worker','34567off'),
('23412', 'Dev', 'Malhotra', 'Bajaj', 'Male', '1983-01-23', '7654321098', 'Belagavi', 7, 'Poll Booth Worker','45678off'),
('34123', 'Saanvi', 'Khanna', 'Naidu', 'Female', '1977-12-07', '6543210987', 'Bangalore Rural', 5, 'Poll Booth Worker','56789off'),
('41234', 'Advik', 'Iyer', 'Verma', 'Male', '1986-05-30', '5432109876', 'Mandya', 9, 'Poll Booth Worker','67890off'),
('52341', 'Tara', 'Varma', 'Dubey', 'Female', '1993-08-11', '4321098765', 'Bangalore Urban', 3, 'Poll Booth Worker','34567off'),
('63412', 'Aditi', 'Sharma', 'Kapoor', 'Female', '1974-11-26', '3210987654', 'Belagavi', 6, 'Poll Booth Worker','45678off'),
('74522', 'Virat', 'Mehra', 'Jain', 'Male', '1981-02-19', '2109876543', 'Bangalore Rural', 4, 'Poll Booth Worker','56789off'),
('85634', 'Aisha', 'Khurana', 'Patel', 'Female', '1989-07-14', '1098765432', 'Mandya', 8, 'Poll Booth Worker','67890off'),
('96745', 'Kabir', 'Das', 'Singh', 'Male', '1978-06-02', '9876543210', 'Bangalore Urban', 1, 'Poll Booth Worker','34567off'),
('07856', 'Zara', 'Sengupta', 'Chopra', 'Female', '1985-09-08', '8765432109', 'Belagavi', 7, 'Poll Booth Worker','45678off'),
('18967', 'Rohan', 'Menon', 'Chaudhary', 'Male', '1991-04-15', '7654321098', 'Mandya', 9, 'Poll Booth Worker','67890off'),
('29078', 'Isha', 'Dutta', 'Roy', 'Female', '1986-11-23', '6543210987', 'Bangalore Rural', 5, 'Poll Booth Worker','56789off'),
('30189', 'Aaradhya', 'Sinha', 'Sharma', 'Female', '1980-08-02', '5432109876', 'Bangalore Urban', 2, 'Poll Booth Worker','34567off'),
('41290', 'Aadi', 'Rao', 'Chopra', 'Male', '1983-12-30', '4321098765', 'Belagavi', 6, 'Poll Booth Worker','45678off'),
('52301', 'Misha', 'Kapoor', 'Sengupta', 'Female', '1994-01-27', '3210987654', 'Mandya', 8, 'Poll Booth Worker','67890off'),
('63410', 'Vihaan', 'Gupta', 'Malhotra', 'Male', '1975-06-19', '2109876543', 'Bangalore Rural', 4, 'Poll Booth Worker','56789off'),
('74523', 'Myra', 'Singh', 'Chadha', 'Female', '1988-10-09', '1098765432', 'Bangalore Urban', 3, 'Poll Booth Worker','34567off'),
('85633', 'Yash', 'Shah', 'Chopra', 'Male', '1993-03-04', '9876543210', 'Belagavi', 7, 'Poll Booth Worker','45678off'),
('96744', 'Anika', 'Chatterjee', 'Gupta', 'Female', '1977-07-22', '8765432109', 'Mandya', 9, 'Poll Booth Worker','67890off'),
('07855', 'Reyansh', 'Kumar', 'Sharma', 'Male', '1984-09-14', '7654321098', 'Bangalore Rural', 5, 'Poll Booth Worker','56789off'),
('18966', 'Avni', 'Rai', 'Sinha', 'Female', '1981-05-31', '6543210987', 'Bangalore Urban', 1, 'Poll Booth Worker','34567off'),
('29079', 'Aahana', 'Dube', 'Gupta', 'Female', '1976-01-10', '5432109876', 'Bangalore Rural', 4, 'Poll Booth Worker','56789off'),
('30180', 'Rudra', 'Rao', 'Chadha', 'Male', '1989-08-25', '4321098765', 'Belagavi', 6, 'Poll Booth Worker','45678off'),
('41292', 'Aarya', 'Menon', 'Sengupta', 'Female', '1974-04-03', '3210987654', 'Mandya', 8, 'Poll Booth Worker','67890off');

alter table candidate add constraint uniquecandcons unique (cons_fight,party_rep);

delimiter //
create procedure getconsinfo(in input_voter_id varchar(20)) begin select c.constituency_name, c.state, c.voter_count, concat(coalesce(cand.first_name,''),' ',coalesce(cand.middle_name,''),' ',coalesce(cand.last_name,'')) as candidate_name, cand.age,p.party_name,p.party_symbol from constituency c left join candidate cand on cand.cons_fight = c.constituency_name left join party p on cand.party_rep = p.party_name where cand.cons_fight in (select constituency_name from voter where voter_id = input_voter_id); end//
delimiter ;

delimiter //
create function malecount(cname varchar(30)) returns int deterministic begin declare male_count int; select sum(case when gender = 'Male' then 1 else 0 end) into male_count from voter where constituency_name = cname; return male_count; end//

create function femalecount(cname varchar(30)) returns int deterministic begin declare female_count int; select sum(case when gender = 'Female' then 1 else 0 end) into female_count from voter where constituency_name = cname; return female_count; end//

create procedure getconsdets() begin select c.constituency_name, malecount(c.constituency_name) as male_count, femalecount(c.constituency_name) AS female_count, count(distinct pb.poll_booth_id) as poll_booth_count from constituency c left join poll_booth pb on c.constituency_id = pb.constituency_id group by c.constituency_name; end//
delimiter ;

