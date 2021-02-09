#include<iostream>
using namespace std;

struct node {
    int data;
    node *next;
};

class Queue{
    node *front, *rear;

    public:
        Queue(): front(NULL), rear(NULL){}
        void insert(int);
        void del();
        void display();
};

void Queue :: insert(int value){
    node *new_node = new node;
    new_node->data = value;
    new_node->next = NULL;

    if(rear == NULL){
        front = new_node;
        rear = new_node;
    }else{
        rear->next = new_node;
        rear = new_node;
    }
}

void Queue :: del(){
    if(front==NULL){
        cout<<"The Queue is empty !\n";
    }else{
        node *ptr = front;
        front = front->next;
        cout<<"The deleted node is: "<<ptr->data<<endl;
        delete ptr;
    }

    if(front==NULL){
        rear = NULL;
    }
}

void Queue::display(){
    if(front==NULL){
        cout<<"The Queue is empty !\n";
    }else{
        node *ptr = front;

        while(ptr != NULL){
            cout<<ptr->data<<" ";
            ptr = ptr->next;
        }
        cout<<endl;
    }
}

int main(){
    int option, value;
    Queue q;

    do{
        cout<<"\n************ MENU *************\n"
            <<"1. Insert\n"
            <<"2. Delete\n"
            <<"3. Display\n"
            <<"4. Exit\n";
        
        cout<<"Enter your option: ";
        cin>>option;

        switch(option){
            case 1:
                cout<<"Enter the value to be inserted: ";
                cin>>value;
                q.insert(value);
                break;
            case 2:
                q.del();
                break;
            case 3:
                q.display();
                break;
            case 4:
                break;
            default:
                cout<<"Wrong option !\n";
                break;
        }
    }while(option != 4);

    if(option==4){
        cout<<"\nThank you for using the program !\n"
            <<"Have a good day !\n";
    }
}