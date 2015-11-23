class Person < ActiveRecord::Base
end

class Debt < ActiveRecord::Base
end

class Expense < ActiveRecord::Base
end

["people", "debts", "expenses"].each do |table_name|
  ActiveRecord::Base.connection.execute("TRUNCATE #{table_name} RESTART IDENTITY")
end

House.create(:name => "Test House")

Person.create(:name => "Jared McFarland", :email => "jared.online@gmail.com", :house_id => 1)
Person.create(:name => "Phillip McFarland", :email => "pmcfarla@email.arizona.edu", :house_id => 1)

Debt.create(:amount_in_cents => 500, :person_id => 1, :loaner_id => 1, :expense_id => 1)
Debt.create(:amount_in_cents => 500, :person_id => 2, :loaner_id => 1, :expense_id => 1)

Expense.create(:amount_in_cents => 1000, :house_id => 1, :name => "Test expense", :description => "Seeded expense", :notes => "la la la", :loaner_id => 1)
