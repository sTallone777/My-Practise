#include<iostream>
#include<vector>
#include<cstring>

using namespace std;

#define STARTING_SIZE 101
#define EXPAND_FACTOR 0.7

class HashTable{
    private:
        int elements;
        int table_size;
        int *hash_table;

        int hash(int value){
            return (value % this->table_size);
        }

        int find_position(int value){
            int hash = this->hash(value);
            for(int i=0; i<this->table_size; i++){
                if(this->hash_table[hash] == value){
                    return hash;
                }
                if(this->hash_table[hash] == 0){
                    return hash;
                }
                if(hash++ == this->table_size){
                    hash = 0;
                }
            }
            return -1;
        }

        void expand(){
            int *new_table = new int[this->table_size * 2];
            if(!new_table){
                cout<<"The HashTable can't be expanded !"<<endl;
                return;
            }

            int *old_talbe = this->hash_table;
            int old_table_size = this->table_size;
            this->hash_table = new_table;
            this->table_size = this->table_size * 2;
            memset(this->hash_table, 0, sizeof(int) * STARTING_SIZE);

            for(int i=0; i<old_table_size; i++){
                if(old_talbe[i] != 0){
                    this->insert(old_talbe[i]);
                }
            }
            delete[] old_talbe;
        }

    public:
        HashTable(){
            this->elements = 0;
            this->table_size = STARTING_SIZE;
            this->hash_table = new int[STARTING_SIZE];
            memset(this->hash_table, 0, sizeof(int) * STARTING_SIZE);
        }

        ~HashTable(){
            if(this->hash_table){
                delete[] this->hash_table;
            }
        }

        void insert(int value){
            int position = this->find_position(value);
            if(position == -1){
                cout<<"The table is full and can't be expanded !"<<endl;
                return;
            }
            if(this->hash_table[position] != value){
                this->hash_table[position] = value;
                this->elements++;
            }
            if((float)(this->elements / this->table_size) > EXPAND_FACTOR){
                this->expand();
            }
        }

        bool find(int value){
            int position = this->find_position(value);
            if(position == -1){
                return false;
            }
            if(this->hash_table[position] == value){
                return true;
            }
            return false;
        }

        bool del(int value){
            int position = this->find_position(value);
            if (this->hash_table[position] == value){
                this->hash_table[position] = 0;
                this->elements--;
                return true;
            }
            return false;
        }

        void display(){
            for(int i=0; i<this->table_size; i++){
                if(hash_table[i] != 0){
                    cout<<hash_table[i]<< " ";
                }
            }
            cout<<endl;
        }
};

int main(){
    int option, value;
    HashTable hashtable;

    do{
        cout<<"\n*********** MENU ***********\n"
            <<"1. Insert\n"
            <<"2. Search\n"
            <<"3. Delete\n"
            <<"4. Display\n"
            <<"5. Exit\n";
        
        cout<<"Enter your option: ";
        cin>>option;

        switch(option){
            case 1: cout<<"Enter the value to be inserted: ";
            cin>>value;
            hashtable.insert(value);
            break;

            case2: cout<<"Enter the value to be searched: ";
            cin>>value;

            if(hashtable.find(value)){
                cout<<"The element was found !"<<endl;
            }else{
                cout<<"The element was not found !"<<endl;
            }
            break;
            case 3: cout<<"Enter the value to be deleted: ";
            cin>>value;

            if(hashtable.del(value)){
                cout<<"The element was deleted !"<<endl;
            }else{
                cout<<"The element was not found !"<<endl;
            }
            break;

            case 4: hashtable.display();
            break;

            case 5: break;

            default: cout<<"Wrong option !\n";
            break;
        }
    }while(option != 5);

    if(option == 5){
        cout<<"\hThank you for using the program !\h"
            <<"Have a good day !\n";
    }

    system("pause");
}