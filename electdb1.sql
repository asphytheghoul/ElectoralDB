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

create table campaign(campaign_id int primary key, expected_crowd int, cons_loc varchar(30), campaign_date date, party_name char(25), foreign key (cons_loc) references constituency(constituency_name), foreign key (party_name) references party(party_name));

create table candidate(aadhar_id int unique, first_name varchar(30) NOT NULL, last_name varchar(30) NOT NULL, middle_name varchar(30) default NULL, gender char(6) NOT NULL, dob date NOT NULL, age int, phone_no int NOT NULL,cons_fight varchar(30), candidate_id varchar(20) primary key, foreign key (cons_fight) references constituency(constituency_name));

delimiter //
create trigger set_canidate_id before insert on candidate for each row begin set NEW.candidate_id = CONCAT(NEW.aadhar_id,'off'); end// 
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